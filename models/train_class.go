package models

import (
	"../connection"
	"log"
	"time"
)

type TrainClass struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	TrainID   int    `gorm:"INTEGER; NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&TrainClass{})

	log.Println("Initialize Train Class Success")
}

func GetAllTrainClass() []TrainClass {
	database := connection.GetConnection()
	defer database.Close()

	var trainClass []TrainClass
	database.Find(&trainClass)

	return trainClass
}

func SearchTrainClassByName(name string) TrainClass {
	database := connection.GetConnection()
	defer database.Close()

	var trainClass TrainClass
	database.
		Where("name = ?", name).
		First(&trainClass)

	return trainClass
}

func InsertTrainClass(name string, id int) *TrainClass {
	database := connection.GetConnection()
	defer database.Close()

	newTrainClass := &TrainClass{
		Name:      name,
		TrainID:   id,
	}
	database.Save(newTrainClass)

	log.Println("Insert New Train Class Success")
	return newTrainClass
}
