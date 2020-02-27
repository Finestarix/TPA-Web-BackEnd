package hotel

import (
	"../../connection"
	"log"
	"time"
)

type HotelFacility struct {
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
		AutoMigrate(&HotelFacility{})

	log.Println("Initialize Hotel Facility Success")
}

func GetAllHotelFacility() []HotelFacility {
	database := connection.GetConnection()
	defer database.Close()

	var hotelFacility []HotelFacility
	database.Find(&hotelFacility)

	return hotelFacility
}

func InsertHotelFacility(hotelID int, facilityName string) *HotelFacility {
	database := connection.GetConnection()
	defer database.Close()

	imageName := facilityName + ".png"

	newHotelFacility := &HotelFacility{
		Name:      facilityName,
		HotelID:   hotelID,
		Photo:     imageName,
	}
	database.Save(newHotelFacility)

	log.Println("Insert New Hotel Facility Success")
	return newHotelFacility
}



