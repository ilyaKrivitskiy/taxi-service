package repository

import "github.com/jmoiron/sqlx"

type UserAuthorization interface {
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
	return &Repository{}
}
