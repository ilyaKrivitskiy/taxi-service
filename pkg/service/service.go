package service

import "github.com/ilyaKrivitskiy/taxi-service/pkg/repository"

type UserAuthorization interface {
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
	return &Service{}
}
