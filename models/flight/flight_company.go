package flight

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type FlightCompany struct {
	ID        int    `gorm:"PRIMARY_KEY"`
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
		AutoMigrate(&FlightCompany{})

	log.Println("Initialize Flight Company Success")
}

func GetAllFlightCompany() []FlightCompany {
	database := connection.GetConnection()
	defer database.Close()

	var company []FlightCompany
	if core.ValidateAPIKey() == false {
		return company
	}
	database.Find(&company)

	return company
}

func GetFlightCompanyByName(name string) FlightCompany {
	database := connection.GetConnection()
	defer database.Close()

	var company FlightCompany
	if core.ValidateAPIKey() == false {
		return company
	}
	database.Where("name = ?", name).First(&company)

	return company
}

func InsertFlightCompany(name string, image string) *FlightCompany {
	database := connection.GetConnection()
	defer database.Close()

	newCompany := &FlightCompany{
		Name:  name,
		Image: image,
	}
	database.Save(newCompany)

	log.Println("Insert New Flight Company Success")
	return newCompany
}
