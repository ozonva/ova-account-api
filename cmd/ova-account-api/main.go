package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonva/ova-account-api/internal/api"
	"github.com/ozonva/ova-account-api/internal/app"
	"github.com/rs/zerolog"
)

func main() {
	confPath := flag.String("conf", "./configs/app.json", "Location of the configuration file")
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	application, err := app.Init(*confPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer func() {
		if err := application.Release(); err != nil {
			logger.Error().Err(err)
		}
	}()

	fmt.Printf("Service %s.\n", application.Conf.Name)

	updatingConfiguration(logger, application.Conf, *confPath, 10)

	logger.Info().Msg("Starting the server...")
	server := api.NewServer(logger, application)
	server.Start()
	logger.Info().Msgf("The application is ready to serve requests on port %s.", application.Conf.GrpcPort)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		logger.Info().Msgf("Received a signal: %v", x)
	case err := <-server.Notify():
		logger.Error().Err(err).Msg("Received an error from the grpc server")
	}

	logger.Info().Msg("Stopping the server...")
	server.Stop()
}

func updatingConfiguration(logger zerolog.Logger, config *app.Config, path string, times int) {
	for i := 0; i < times; i++ {
		err := config.Update(path)
		if err != nil {
			logger.Err(err).Msg("Unable to update the configuration")
			continue
		}

		logger.Info().Msg("Successfully updated the configuration")
	}
}
