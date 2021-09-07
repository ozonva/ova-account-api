package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonva/ova-account-api/internal/api"
	"github.com/ozonva/ova-account-api/internal/app"
	"github.com/ozonva/ova-account-api/internal/metrics"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func main() {
	confPath := flag.StringP("conf", "c", ".env", "location of the configuration file")
	_ = flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	application, err := app.Init(*confPath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() {
		if err := application.Release(); err != nil {
			log.Error().Err(err).Msg("Got an error while releasing app resources.")
		}
	}()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if application.Conf.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	fmt.Printf("Service %s.\n", application.Conf.Name)

	updatingConfiguration(application.Conf, 10)

	log.Info().Msg("Starting the servers...")
	server := api.NewServer(application)
	server.Start()

	metricsServer := metrics.NewServer()
	metricsServer.Start()
	log.Info().Msgf("The application is ready to serve requests on port %s.", application.Conf.GrpcPort)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Info().Msgf("Received a signal: %v", x)
	case err := <-server.Notify():
		log.Error().Err(err).Msg("Received an error from the grpc server")
	case err := <-metricsServer.Notify():
		log.Error().Err(err).Msg("Received an error from the metrics server")
	}

	log.Info().Msg("Stopping the servers...")
	server.Stop()

	if err := metricsServer.Stop(); err != nil {
		log.Error().Err(err).Msg("Got an error while stopping the metrics server.")
	}
}

func updatingConfiguration(config *app.Config, times int) {
	for i := 0; i < times; i++ {
		err := config.Update("./configs/app.json")
		if err != nil {
			log.Err(err).Msg("Unable to update the configuration")
			continue
		}

		log.Info().Msg("Successfully updated the configuration")
	}
}
