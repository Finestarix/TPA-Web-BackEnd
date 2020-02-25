package models

import (
	"../connection"
	"log"
	"time"
)

type TrainStation struct {
	ID         int      `gorm:"PRIMARY_KEY"`
	Name       string   `gorm:"VARCHAR(100); NOT NULL"`
	Code       string   `gorm:"VARCHAR(100); NOT NULL"`
	Location   Location `gorm:"FOREIGNKEY:LocationID"`
	LocationID int      `gorm:"INTEGER; NOT NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&TrainStation{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Train Station Success")
}

func GetAllTrainStation() []TrainStation {
	database := connection.GetConnection()
	defer database.Close()

	var trainStation []TrainStation
	database.Find(&trainStation)

	for i, _ := range trainStation {
		database.Model(trainStation[i]).Related(&trainStation[i].Location, "location_id")
	}

	return trainStation
}

func InsertTrainStation(name string, code string, city string) *TrainStation {
	database := connection.GetConnection()
	defer database.Close()

	location := GetLocationByCity(city)

	newTrainStation := &TrainStation{
		Name:       name,
		Code:       code,
		LocationID: location.ID,
	}
	database.Save(newTrainStation)

	database.Model(newTrainStation).Related(&newTrainStation.Location, "location_id")

	log.Println("Insert New Train Station Success")
	return newTrainStation
}

func SearchTrainStationByName(name string) TrainStation {
	database := connection.GetConnection()
	defer database.Close()

	var trainStation TrainStation
	database.
		Where("name = ?", name).
		First(&trainStation)

	database.Model(trainStation).Related(&trainStation.Location, "location_id")

	return trainStation
}
