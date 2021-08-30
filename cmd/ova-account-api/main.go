package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonva/ova-account-api/internal/api"
	"github.com/ozonva/ova-account-api/internal/app"
)

func main() {
	confPath := flag.String("conf", "./configs/app.json", "Location of the configuration file")
	flag.Parse()

	config, err := app.ParseConfig(*confPath)
	if err != nil {
		log.Fatal("Can't process the configuration: ", err)
	}

	fmt.Printf("Service %s.\n", config.Name)

	updatingConfiguration(config, *confPath, 10)

	log.Println("Starting the server...")
	server := api.NewServer(config.GrpcPort)
	server.Start()
	log.Printf("The application is ready to serve requests on port %s.", config.GrpcPort)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Println("Received a signal:", x)
	case err := <- server.Notify():
		log.Println("Received an error from the grpc server:", err)
	}

	log.Println("Stopping the server...")
	server.Stop()
}

func updatingConfiguration(config *app.Config, path string, times int) {
	for i := 0; i < times; i++ {
		err := config.Update(path)
		if err != nil {
			log.Println("Unable to update the configuration:", err)
			continue
		}

		log.Println("Successfully updated the configuration")
	}
}
