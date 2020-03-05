package flight

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type Flight struct {
	ID                int           `gorm:"PRIMARY_KEY"`
	Company           FlightCompany `gorm:"FOREIGNKEY:CompanyID"`
	Model             string        `gorm:"VARCHAR(100); NOT NULL;"`
	CompanyID         int           `gorm:"INTEGER; NOT NULL;"`
	FromAirport       FlightAirport `gorm:"FOREIGNKEY:FromAirportID"`
	FromAirportID     int           `gorm:"INTEGER"`
	FromLocationID    int           `gorm:"INTEGER"`
	ArrivalTime       time.Time
	ToAirport         FlightAirport `gorm:"FOREIGNKEY:ToAirportID"`
	ToAirportID       int           `gorm:"INTEGER"`
	ToLocationID      int           `gorm:"INTEGER"`
	DepartureTime     time.Time
	Price             int              `gorm:"INTEGER"`
	Facility          []FlightFacility `gorm:"FOREIGNKEY:FlightID"`
	Transit           FlightAirport    `gorm:"FOREIGNKEY:TransitID"`
	TransitID         int              `gorm:"INTEGER"`
	TransitDuration   int
	TransitLocationID int `gorm:"INTEGER"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Flight{}).
		AddForeignKey("company_id", "flight_companies(id)", "CASCADE", "CASCADE").
		AddForeignKey("from_airport_id", "flight_airports(id)", "CASCADE", "CASCADE").
		AddForeignKey("to_airport_id", "flight_airports(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Flight Airport Success")
}

func GetAllFlight() []Flight {
	database := connection.GetConnection()
	defer database.Close()

	var flight []Flight
	if core.ValidateAPIKey() == false {
		return flight
	}
	database.Find(&flight)

	for i, _ := range flight {
		database.Model(&flight[i]).Related(&flight[i].Company, "company_id")
		database.Model(&flight[i]).Related(&flight[i].FromAirport, "from_airport_id")
		database.Model(&flight[i]).Related(&flight[i].FromAirport.Location, "from_location_id")
		database.Model(&flight[i]).Related(&flight[i].ToAirport, "to_airport_id")
		database.Model(&flight[i]).Related(&flight[i].ToAirport.Location, "to_location_id")
		database.Model(&flight[i]).Related(&flight[i].Facility, "flight_id")
		database.Model(&flight[i]).Related(&flight[i].Transit, "transit_id")
		database.Model(&flight[i]).Related(&flight[i].Transit.Location, "transit_location_id")
	}

	return flight
}

func GetFlightByID(id int) Flight {
	database := connection.GetConnection()
	defer database.Close()

	var flight Flight
	if core.ValidateAPIKey() == false {
		return flight
	}
	database.
		Where("id = ?", id).
		First(&flight)

	return flight
}

func GetFlightByAirport(fromAirportName string, toAirportName string, date time.Time) []Flight {
	database := connection.GetConnection()
	defer database.Close()

	fromAirport := GetFlightAirportByCode(fromAirportName)
	toAirport := GetFlightAirportByCode(toAirportName)

	var flight []Flight
	if core.ValidateAPIKey() == false {
		return flight
	}
	database.
		Where("from_airport_id = ? AND to_airport_id = ? AND DATE_PART('day',arrival_time) = ?",
			fromAirport.ID, toAirport.ID, date.Day()).
		Find(&flight)

	for i, _ := range flight {
		database.Model(&flight[i]).Related(&flight[i].Company, "company_id")
		database.Model(&flight[i]).Related(&flight[i].FromAirport, "from_airport_id")
		database.Model(&flight[i]).Related(&flight[i].FromAirport.Location, "from_location_id")
		database.Model(&flight[i]).Related(&flight[i].ToAirport, "to_airport_id")
		database.Model(&flight[i]).Related(&flight[i].ToAirport.Location, "to_location_id")
		database.Model(&flight[i]).Related(&flight[i].Facility, "flight_id")
		database.Model(&flight[i]).Related(&flight[i].Transit, "transit_id")
		database.Model(&flight[i]).Related(&flight[i].Transit.Location, "transit_location_id")
	}
	return flight
}

func InsertFlight(companyName string, fromAirportName string, arrivalTime time.Time,
	toAirportName string, departureTime time.Time, price int, model string, transitDuration int,
	transitAirportName string) *Flight {

	database := connection.GetConnection()
	defer database.Close()

	if price <= 0 {
		return nil
	}

	company := GetFlightCompanyByName(companyName)
	fromAirport := GetFlightAirportByCode(fromAirportName)
	toAirport := GetFlightAirportByCode(toAirportName)
	transitID := 1000
	transitLocationID := 1
	var transitAirport FlightAirport
	if len(transitAirportName) != 0 {
		transitAirport = GetFlightAirportByCode(transitAirportName)
		transitID = transitAirport.ID
		transitLocationID = transitAirport.LocationID
	}

	newFlight := &Flight{
		CompanyID:         company.ID,
		FromAirportID:     fromAirport.ID,
		ArrivalTime:       arrivalTime.Add(-7 * time.Hour),
		ToAirportID:       toAirport.ID,
		DepartureTime:     departureTime.Add(-7 * time.Hour),
		TransitID:         transitID,
		Price:             price,
		Model:             model,
		TransitDuration:   transitDuration,
		FromLocationID:    fromAirport.LocationID,
		ToLocationID:      toAirport.LocationID,
		TransitLocationID: transitLocationID,
	}
	database.Save(newFlight)

	log.Println("Insert New Flight Success")
	return newFlight
}

func UpdateFlight(id int, fromAirportName string, arrivalTime time.Time, transitDuration int,
	toAirportName string, departureTime time.Time, price int, model string, transitAirportName string) Flight {
	database := connection.GetConnection()
	defer database.Close()

	fromAirport := GetFlightAirportByCode(fromAirportName)
	toAirport := GetFlightAirportByCode(toAirportName)
	transitID := 1000
	transitLocationID := 1
	var transitAirport FlightAirport
	if len(transitAirportName) != 0 {
		transitAirport = GetFlightAirportByCode(transitAirportName)
		transitID = transitAirport.ID
		transitLocationID = transitAirport.LocationID
	}

	var updateFlight Flight
	database.
		Model(&updateFlight).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"FromAirportID":     fromAirport.ID,
			"TransitDuration":   transitDuration,
			"ArrivalTime":       arrivalTime.Add(-7 * time.Hour),
			"ToAirportID":       toAirport.ID,
			"DepartureTime":     departureTime.Add(-7 * time.Hour),
			"TransitID":         transitID,
			"Price":             price,
			"Model":             model,
			"FromLocationID":    fromAirport.LocationID,
			"ToLocationID":      toAirport.LocationID,
			"TransitLocationID": transitLocationID,
		})

	database.Model(&updateFlight).Related(&updateFlight.Company, "company_id")
	database.Model(&updateFlight).Related(&updateFlight.FromAirport, "from_airport_id")
	database.Model(&updateFlight).Related(&updateFlight.FromAirport.Location, "from_location_id")
	database.Model(&updateFlight).Related(&updateFlight.ToAirport, "to_airport_id")
	database.Model(&updateFlight).Related(&updateFlight.ToAirport.Location, "to_location_id")
	database.Model(&updateFlight).Related(&updateFlight.Facility, "updateFlight_id")
	database.Model(&updateFlight).Related(&updateFlight.Transit, "transit_id")
	database.Model(&updateFlight).Related(&updateFlight.Transit.Location, "transit_location_id")

	log.Println("Insert New Flight Success")
	return updateFlight
}

func DeleteFlight(id int) *Flight {
	database := connection.GetConnection()
	defer database.Close()

	var flight Flight
	flight = GetFlightByID(id)

	if flight.ID == 0 {
		log.Println("Delete Flight Failed")
		return &flight
	}

	err := database.Delete(flight).Error

	if err != nil {
		panic("Error Delete Hotel !" + err.Error())
	}

	log.Println("Delete Flight Success")
	return &flight
}
