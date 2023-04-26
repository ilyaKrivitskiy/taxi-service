package models

type Order struct {
	Id        uint   `json:"-"`
	UserId    uint   `json:"user_id"`
	DriverId  uint   `json:"driver_id"`
	PlaceFrom string `json:"from"`
	PlaceTo   string `json:"to"`
	TaxiType  string `json:"taxi_type"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
	Status    string `json:"status"`
	Info      string `json:"comment"`
}
