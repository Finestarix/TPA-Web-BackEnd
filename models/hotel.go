package models

import (
	"../connection"
	"log"
	"math"
	"time"
)

type Hotel struct {
	ID         int            `gorm:"PRIMARY_KEY"`
	Name       string         `gorm:"VARCHAR(100); NOT NULL"`
	Address    string         `gorm:"VARCHAR(100); NOT NULL"`
	Location   Location       `gorm:"FOREIGNKEY:LocationID"`
	LocationID int            `gorm:"INTEGER; NOT NULL"`
	Longitude  float64        `gorm:"DECIMAL(3,1)"`
	Latitude   float64        `gorm:"DECIMAL(3,1)"`
	Photo      []HotelImage   `gorm:"FOREIGNKEY:HotelID"`
	Facility   []HotelFacility `gorm:FOREIGNKEY:HotelID`
	Price      int            `gorm:"INTEGER; NOT NULL"`
	Rating     float64        `gorm:"DECIMAL(3,1)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Hotel{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Hotel Success")
}

func DropTableHotel() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&Hotel{})
	database.AutoMigrate(&Hotel{}).AddForeignKey("location_id", "location(id)", "CASCADE", "CASCADE")

	log.Println("Drop Database Success")
}

func GetAllHotel() []Hotel {
	database := connection.GetConnection()
	defer database.Close()

	var hotels []Hotel
	database.Find(&hotels)

	for i, _ := range hotels {
		database.Model(hotels[i]).Related(&hotels[i].Location, "location_id")
		database.Model(hotels[i]).Related(&hotels[i].Photo, "HotelID")
		database.Model(hotels[i]).Related(&hotels[i].Facility, "HotelID")
	}

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

func InsertHotel(name string, address string, city string, price int, rating float64, latitude float64, longitude float64) *Hotel {
	database := connection.GetConnection()
	defer database.Close()

	location := GetLocationByCity(city)

	newHotel := &Hotel{
		Name:       name,
		Address:    address,
		LocationID: location.ID,
		Price:      price,
		Rating:     rating,
		Latitude:   latitude,
		Longitude:  longitude,
	}
	database.Save(newHotel)

	log.Println("Insert New Hotel Success")
	return newHotel
}

func distance(currLatitude float64, currLongitude float64, searchLatitude float64, searchLongitude float64) float64 {
	const PI float64 = 3.141592653589793

	radianLatitude1 := float64(PI * currLatitude / 180)
	radianLatitude2 := float64(PI * searchLatitude / 180)

	theta := float64(currLongitude - searchLongitude)
	radianTheta := float64(PI * theta / 180)

	distance := math.Sin(radianLatitude1)*math.Sin(radianLatitude2) +
		math.Cos(radianLatitude1)*math.Cos(radianLatitude2)*
			math.Cos(radianTheta)

	if distance > 1 {
		distance = 1
	}

	distance = math.Acos(distance)
	distance = distance * 180 / PI
	distance = distance * 60 * 1.1515

	return distance * 1.609344
}

func compareDistance(currLatitude float64, currLongitude float64, hotel1 Hotel, hotel2 Hotel) bool {

	distance1 := distance(currLatitude, currLongitude, hotel1.Latitude, hotel1.Longitude)
	distance2 := distance(currLatitude, currLongitude, hotel2.Latitude, hotel2.Longitude)

	if distance1 < distance2 {
		return true
	} else {
		return false
	}

}

func bubbleSort(currLatitude float64, currLongitude float64, hotels []Hotel) {
	size := len(hotels)

	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if compareDistance(currLatitude, currLongitude, hotels[j], hotels[j-1]) {
				hotels[j], hotels[j-1] = hotels[j-1], hotels[j]
			}
		}
	}
}

func GetNearestHotel(currLatitude float64, currLongitude float64) []Hotel {
	db := connection.GetConnection()
	defer db.Close()

	var hotels []Hotel

	// With Query (Failed)
	//db.Order("((ACOS(SIN(latitude * PI() / 180) * SIN(latitude * PI() / 180) + COS(latitude * PI() / 180) * COS(latitude * PI() / 180) * COS((longitude - longitude) * PI() / 180)) * 180 / PI()) * 60 * 1.1515)").Find(&hotels)

	db.Find(&hotels)
	for i, _ := range hotels {
		db.Model(hotels[i]).Related(&hotels[i].Location, "location_id")
		db.Model(hotels[i]).Related(&hotels[i].Photo, "HotelID")
		db.Model(hotels[i]).Related(&hotels[i].Facility, "HotelID")
	}

	//fmt.Println("Before: ")
	//for i, _ := range hotels {
	//	fmt.Printf("Curr to %s: %f KM\n", hotels[i].Location.City, distance(currLatitude, currLongitude, hotels[i].Latitude, hotels[i].Longitude))
	//}

	bubbleSort(currLatitude, currLongitude, hotels)

	//fmt.Println("After: ")
	//for i, _ := range hotels {
	//	fmt.Printf("Curr to %s: %f KM\n", hotels[i].Location.City, distance(currLatitude, currLongitude, hotels[i].Latitude, hotels[i].Longitude))
	//}

	hotels = hotels[0:8]

	return hotels
}

func GetHotelByProvince(province string) []Hotel {
	database := connection.GetConnection()
	defer database.Close()

	location := GetLocationByProvince(province)

	var hotels []Hotel

	if len(location) == 1 {
		database.Where("location_id = ?", location[0].ID).Find(&hotels)
	} else {
		var listLocationID []int
		for i := 0 ; i < len(location) ; i++ {
			listLocationID = append(listLocationID, location[i].ID)
		}
		database.Where("location_id IN (?)", listLocationID).Find(&hotels)
	}

	for i, _ := range hotels {
		database.Model(hotels[i]).Related(&hotels[i].Location, "location_id")
		database.Model(hotels[i]).Related(&hotels[i].Photo, "HotelID")
		database.Model(hotels[i]).Related(&hotels[i].Facility, "HotelID")
	}

	return hotels
}

func GetHotelByLatLong(latitude float64, longitude float64) Hotel {
	database := connection.GetConnection()
	defer database.Close()

	var hotel Hotel
	database.
		Where("latitude = ?", latitude).
		Where("longitude = ?", longitude).
		First(&hotel)

	database.Model(hotel).Related(&hotel.Location, "location_id")
	database.Model(hotel).Related(&hotel.Photo, "HotelID")
	database.Model(hotel).Related(&hotel.Facility, "HotelID")

	return hotel
}