package models

type Order struct {
	UserId   uint   `json:"user_id"`
	DriverId uint   `json:"driver_id"`
	From     string `json:"from"`
	To       string `json:"to"`
	TaxiType string `json:"taxi_type"`
	Date     string `json:"date"`
	Status   string `json:"status"`
	Comment  string `json:"comment"`
}
