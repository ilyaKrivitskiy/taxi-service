package models

type Driver struct {
	Id       uint   `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	TaxiType string `json:"taxi_type"`
}
