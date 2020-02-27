package car

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type Car struct {
	ID         int           `gorm:"PRIMARY_KEY"`
	CarModel   CarModel      `gorm:"FOREIGNKEY:CarModelID"`
	CarModelID int           `gorm:"INTEGER; NOT NULL"`
	Location   core.Location `gorm:"FOREIGNKEY:LocationID"`
	LocationID int           `gorm:"INTEGER; NOT NULL"`
	Price      int           `gorm:"INTEGER; NOT NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Car{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE").
		AddForeignKey("car_model_id", "car_models(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Car Success")
}

func DropTableCar() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&Car{})
	database.AutoMigrate(&Car{})

	log.Println("Drop Database Success")
}

func GetAllCar() []Car {
	database := connection.GetConnection()
	defer database.Close()

	var cars []Car
	database.Find(&cars)

	for i, _ := range cars {
		database.Model(&cars[i]).Related(&cars[i].Location, "location_id")
		database.Model(&cars[i]).Related(&cars[i].CarModel, "car_model_id")
	}

	return cars
}

func GetCarByLocation(city string) []Car {
	database := connection.GetConnection()
	defer database.Close()

	location := core.GetLocationByProvince(city)

	var cars []Car

	if len(location) == 1 {
		database.Where("location_id = ?", location[0].ID).Find(&cars)
	} else {
		var listLocationID []int
		for i := 0; i < len(location); i++ {
			listLocationID = append(listLocationID, location[i].ID)
		}
		database.Where("location_id IN (?)", listLocationID).Find(&cars)
	}

	for i, _ := range cars {
		database.Model(&cars[i]).Related(&cars[i].Location, "location_id")
		database.Model(&cars[i]).Related(&cars[i].CarModel, "car_model_id")
	}

	return cars
}

func InsertCar(model string, city string, price int) *Car {
	database := connection.GetConnection()
	defer database.Close()

	location := core.GetLocationByCity(city)
	carModel := GetCarModelByModel(model)

	newCar := &Car{
		Price:      price,
		CarModelID: carModel.ID,
		LocationID: location.ID,
	}
	database.Save(newCar)

	log.Println("Insert New Car Success")
	return newCar
}
