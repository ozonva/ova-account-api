package app

import (
	"fmt"

	"github.com/ozonva/ova-account-api/internal/repo"
	"github.com/ozonva/ova-account-api/internal/repo/postgres"
)

// App ...
type App struct {
	Conf *Config
	Store repo.Store
}

// Init ...
func Init(configPath string) (*App, error) {
	conf, err := ParseConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("can't process the configuration: %v", err)
	}

	store, err := postgres.NewStore(conf.DB.DSN())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the db: %v", err)
	}

	return &App{
		Conf: conf,
		Store: store,
	}, nil
}

func (a *App) Release() error {
	return a.Store.Close()
}
