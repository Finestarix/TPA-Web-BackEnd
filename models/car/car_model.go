package car

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type CarModel struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Brand     string `gorm:"VARCHAR(100); NOT NULL"`
	Model     string `gorm:"VARCHAR(100); NOT NULL"`
	Passenger int    `gorm:"INTEGER; NOT NULL"`
	Baggage   int    `gorm:"INTEGER; NOT NULL"`
	Image 	  string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func DropTableCarModel() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&CarModel{})
	database.AutoMigrate(&CarModel{})

	log.Println("Drop Database Success")
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&CarModel{})

	log.Println("Initialize Car Model Success")
}

func GetAllCarModel() []CarModel {
	database := connection.GetConnection()
	defer database.Close()

	var cars []CarModel
	if core.ValidateAPIKey() == false {
		return cars
	}
	database.Find(&cars)

	return cars
}

func GetCarModelByModel(model string) CarModel {
	database := connection.GetConnection()
	defer database.Close()

	var cars CarModel
	if core.ValidateAPIKey() == false {
		return cars
	}
	database.
		Where("model = ?", model).
		First(&cars)

	return cars
}

func InsertCarModel(brand string, model string, passenger int, baggage int, image string) *CarModel {
	database := connection.GetConnection()
	defer database.Close()

	newCar := &CarModel{
		Brand:     brand,
		Model:     model,
		Passenger: passenger,
		Baggage:   baggage,
		Image:     image,
	}
	database.Save(newCar)

	log.Println("Insert New Car Success")
	return newCar
}


