package models

import (
	"../connection"
	"log"
	"time"
)

type HotelType struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	HotelID   int    `gorm:"INTEGER"`
	Photo     string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&HotelType{})

	log.Println("Initialize Hotel Type Success")
}

func GetAllHotelType() []HotelType {
	database := connection.GetConnection()
	defer database.Close()

	var hotelType []HotelType
	database.Find(&hotelType)

	return hotelType
}

func InsertHotelType(hotelID int, typeName string) *HotelType {
	database := connection.GetConnection()
	defer database.Close()

	imageName := typeName + ".jpg"

	newHotelType := &HotelType{
		Name:      typeName,
		HotelID:   hotelID,
		Photo:     imageName,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	database.Save(newHotelType)

	log.Println("Insert New Hotel Type Success")
	return newHotelType
}



