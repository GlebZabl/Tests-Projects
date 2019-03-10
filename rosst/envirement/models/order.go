package models

type Order struct {
	Id      string `db:"id"`
	OwnerId string `db:"owner_id"`
	Info    string `db:"info"`
	Date    string `db:"date"`
}
