package flight

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type FlightAirport struct {
	ID         int           `gorm:"PRIMARY_KEY"`
	Name       string        `gorm:"VARCHAR(100); NOT NULL"`
	Code       string        `gorm:"VARCHAR(5); NOT NULL"`
	Location   core.Location `gorm:"FOREIGNKEY:LocationID"`
	LocationID int           `gorm:"INTEGER"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&FlightAirport{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Flight Airport Success")
}

func GetAllFlightAirport() []FlightAirport {
	database := connection.GetConnection()
	defer database.Close()

	var airport []FlightAirport
	database.Where("code != ?", "NULL").Find(&airport)

	for i, _ := range airport {
		database.Model(&airport[i]).Related(&airport[i].Location, "location_id")
	}

	return airport
}

func InsertFlightAirport(name string, code string, city string) *FlightAirport {
	database := connection.GetConnection()
	defer database.Close()

	location := core.GetLocationByCity(city)

	newAirport := &FlightAirport{
		Name:       name,
		Code:       code,
		LocationID: location.ID,
	}
	database.Save(newAirport)

	log.Println("Insert New Flight Airport Success")
	return newAirport
}

func GetFlightAirportByCode(code string) FlightAirport {
	database := connection.GetConnection()
	defer database.Close()

	var airport FlightAirport
	database.Where("code = ?", code).First(&airport)

	database.Model(&airport).Related(&airport.Location, "location_id")

	return airport
}