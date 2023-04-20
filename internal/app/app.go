package app

import (
	"context"
	"github.com/ilyaKrivitskiy/taxi-service/config"
	"github.com/ilyaKrivitskiy/taxi-service/internal/server"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {

	srv := new(server.Server)
	if err := srv.Run(cfg.HTTP.Port); err != nil {
		logrus.Fatalf("Error occured on running server: %s", err)
	}

	logrus.Print("TodoApp Started")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error occured on server shutting down: %s", err)
	}
}
