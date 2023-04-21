package app

import (
	"context"
	"github.com/ilyaKrivitskiy/taxi-service/config"
	"github.com/ilyaKrivitskiy/taxi-service/internal/server"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/handler"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/repository"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/service"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {

	reps := repository.NewRepository()
	servicies := service.NewService(reps)
	handlers := handler.NewHandler(servicies)

	srv := new(server.Server)
	if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured on running server: %s", err)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error occured on server shutting down: %s", err)
	}
}
