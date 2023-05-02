package repository

import (
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"github.com/jmoiron/sqlx"
)

type UserAuthorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type DriverAuthorization interface {
}

type Order interface {
}

type Repository struct {
	UserAuthorization
	DriverAuthorization
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserAuthorization: NewUserAuthPostgres(db),
	}
}
