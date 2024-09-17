package main

import (
	"log"

	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
)

func main() {

	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Config failed to setup environment: %v\n", err)
    }

	api.StartServer(cfg)

}
