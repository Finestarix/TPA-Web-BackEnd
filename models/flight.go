package models

import "time"

type Flight struct {
	ID            int     `gorm:"PRIMARY_KEY"`
	Airline       Airline `gorm:"FOREIGNKEY:AirlineID"`
	AirlineID     int     `gorm:"INTEGER"`
	//Routes        []Route
	FromAirport   Airport `gorm:"FOREIGNKEY:FromAirportID"`
	FromAirportID int     `gorm:"INTEGER"`
	ToAirport     Airport `gorm:"FOREIGNKEY:ToAirportID"`
	ToAirportID   int     `gorm:"INTEGER"`
	DepartureTime time.Time
	ArrivalTime   time.Time
	Duration      int `gorm:"INTEGER"`
	Price         int `gorm:"INTEGER"`
	Tax           int `gorm:"INTEGER"`
	ServiceCharge int `gorm:"INTEGER"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}
