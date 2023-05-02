package service

import (
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/repository"
)

type UserAuthorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type DriverAuthorization interface {
}

type Order interface {
}

type Service struct {
	UserAuthorization
	DriverAuthorization
	Order
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		UserAuthorization: NewUserAuthService(rep),
	}
}
