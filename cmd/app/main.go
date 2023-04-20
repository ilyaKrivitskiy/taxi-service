package main

import (
	"github.com/ilyaKrivitskiy/taxi-service/config"
	"github.com/ilyaKrivitskiy/taxi-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
