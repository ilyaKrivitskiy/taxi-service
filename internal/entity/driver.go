package entity

type Driver struct {
	User
	TaxiType string `json:"taxi_type"`
}
