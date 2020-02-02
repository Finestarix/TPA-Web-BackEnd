package models

import (
	"../connection"
	"log"
	"time"
)

type Hotel struct {
	ID      int    `gorm:"PRIMARY_KEY"`
	Name    string `gorm:"VARCHAR(100); NOT NULL"`
	Address string `gorm:"VARCHAR(100); NOT NULL"`
	City    string `gorm:"VARCHAR(100); NOT NULL"`
	//Photo []string
	Price     int `gorm:"INTEGER; NOT NULL"`
	Rating    float64 `gorm:"DECIMAL(3,1)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.AutoMigrate(&Hotel{})

	log.Println("Initialize Hotel Success")
}

func DropTableHotel() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&Hotel{})
	database.AutoMigrate(&Hotel{})

	log.Println("Drop Database Success")
}

func GetAllHotel() []Hotel {
	database := connection.GetConnection()
	defer database.Close()

	var hotels []Hotel
	database.Find(&hotels)

	return hotels
}

func GetHotelByID(id int) Hotel {
	database := connection.GetConnection()
	defer database.Close()

	var hotel Hotel
	database.
		Where("id = ?", id).
		First(&hotel)

	return hotel
}

func InsertHotel(name string, address string, city string, price int, rating float64) *Hotel {
	database := connection.GetConnection()
	defer database.Close()

	newHotel := &Hotel{
		Name:      name,
		Address:   address,
		City:      city,
		Price:     price,
		Rating:    rating,
	}
	database.Save(newHotel)

	log.Println("Insert New Hotel Success")
	return newHotel
}

