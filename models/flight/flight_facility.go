package flight

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type FlightFacility struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	FlightID  int    `gorm:"INTEGER; NOT NULL"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	Image     string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&FlightFacility{})

	log.Println("Initialize Flight Facility Success")
}

func GetAllFlightFacility() []FlightFacility {
	database := connection.GetConnection()
	defer database.Close()

	var facility []FlightFacility
	if core.ValidateAPIKey() == false {
		return facility
	}
	database.Find(&facility)

	return facility
}

func InsertFlightFacility(flightID int, name string) *FlightFacility {
	database := connection.GetConnection()
	defer database.Close()

	imageLink := name + ".jpg"

	newFacility := &FlightFacility{
		FlightID:  flightID,
		Name:      name,
		Image:     imageLink,
	}
	database.Save(newFacility)

	log.Println("Insert New Flight Facility Success")
	return newFacility
}

