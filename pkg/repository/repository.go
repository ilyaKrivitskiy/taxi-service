package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
