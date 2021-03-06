package app

import (
	"fmt"
	"io"

	"github.com/ozonva/ova-account-api/internal/kafka"
	"github.com/ozonva/ova-account-api/internal/metrics"
	"github.com/ozonva/ova-account-api/internal/repo"
	"github.com/ozonva/ova-account-api/internal/repo/postgres"
)

// App ...
type App struct {
	Conf         *Config
	Store        repo.Store
	Producer     kafka.Producer
	Metrics      metrics.AccountMetrics
	tracerCloser io.Closer
}

// Init ...
func Init(configPath string) (*App, error) {
	conf, err := NewConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("can't process the configuration: %v", err)
	}

	store, err := postgres.NewStore(
		conf.DB.DSN(),
		conf.DB.Pool.MaxOpenConns,
		conf.DB.Pool.MaxIdleConn,
		conf.DB.Pool.ConnMaxLifetime,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the db: %v", err)
	}

	_, closer, err := initTracer(*conf)
	if err != nil {
		return nil, err
	}

	producer := kafka.NewProducer(conf.Kafka.Addr, conf.Kafka.Topic)

	stats := metrics.RegisterMetrics()

	return &App{
		Conf:         conf,
		Store:        store,
		Producer:     producer,
		Metrics:      stats,
		tracerCloser: closer,
	}, nil
}

func (a *App) Release() error {
	_ = a.tracerCloser.Close()
	_ = a.Producer.Close()
	return a.Store.Close()
}
