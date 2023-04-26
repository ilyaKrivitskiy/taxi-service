package app

import (
	"context"
	"github.com/ilyaKrivitskiy/taxi-service/config"
	"github.com/ilyaKrivitskiy/taxi-service/internal/server"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/handler"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/repository"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

func Run(cfg *config.Config) {

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	reps := repository.NewRepository(db)
	servicies := service.NewService(reps)
	handlers := handler.NewHandler(servicies)

	srv := new(server.Server)
	if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured on running server: %s", err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error occured on server shutting down: %s", err.Error())
	}
}
