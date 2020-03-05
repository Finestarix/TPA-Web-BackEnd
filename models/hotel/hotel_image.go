package hotel

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type HotelImage struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	HotelID   int    `gorm:"INTEGER"`
	Source    string `gorm:"VARCHAR(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&HotelImage{}).
		AddForeignKey("hotel_id", "hotels(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Hotel Image Success")
}

func DropTableHotelImage() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&HotelImage{})
	database.
		AutoMigrate(&HotelImage{}).
		AddForeignKey("hotel_id", "hotels(id)", "CASCADE", "CASCADE")

	log.Println("Drop Database Success")
}

func GetAllHotelImage() []HotelImage {
	database := connection.GetConnection()
	defer database.Close()

	var hotelImages []HotelImage
	if core.ValidateAPIKey() == false {
		return hotelImages
	}
	database.Find(&hotelImages)

	return hotelImages
}

func InsertHotelImage(hotelID int, source string) *HotelImage {
	database := connection.GetConnection()
	defer database.Close()

	newHotelImage := &HotelImage{
		HotelID:   hotelID,
		Source:    source,
	}
	database.Save(newHotelImage)

	log.Println("Insert New Hotel Image Success")
	return newHotelImage
}
