package main

import (
	"github.com/ilyaKrivitskiy/taxi-service/config"
	"github.com/ilyaKrivitskiy/taxi-service/internal/app"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Config error: %s", err.Error())
	}

	app.Run(cfg)
}
