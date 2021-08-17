package main

import (
	"flag"
	"fmt"
	"log"

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

	updatingConfiguration(config, *confPath)
}

func updatingConfiguration(config *app.Config, path string) {
	times := 10
	for i := 0; i < times; i++ {
		err := config.Update(path)
		if err != nil {
			log.Println("Unable to update the configuration:", err)
			continue
		}

		log.Println("Successfully updated the configuration")
	}
}
