package repository

import (
	"fmt"
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"github.com/jmoiron/sqlx"
)

type UserAuthPostgres struct {
	db *sqlx.DB
}

func NewUserAuthPostgres(db *sqlx.DB) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

func (r *UserAuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id",
		usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserAuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
