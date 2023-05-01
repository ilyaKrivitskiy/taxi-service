package models

type Driver struct {
	Id       uint   `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	TaxiType string `json:"taxi_type"`
}
