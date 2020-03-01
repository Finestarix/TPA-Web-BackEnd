package event

import (
	"../../connection"
	"log"
	"time"
)

type Entertainment struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Price     int
	Location  string
	Latitude  float64
	Longitude float64
	Date      time.Time
	Type      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Entertainment{})

	log.Println("Initialize Entertainment Success")
}

func GetAllEntertainment() []Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var event []Entertainment
	database.Find(&event)

	return event
}

func GetEntertainmentByID(id int) Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var event Entertainment
	database.Where("id = ?", id).First(&event)

	return event
}

func InsertEntertainment(title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string) *Entertainment {

	database := connection.GetConnection()
	defer database.Close()

	newEvent := &Entertainment{
		Title:     title,
		Price:     price,
		Location:  location,
		Latitude:  latitude,
		Longitude: longitude,
		Date:      date.Add(-7 * time.Hour),
		Type:      category,
		Image:     image,
	}
	database.Save(newEvent)

	log.Println("Insert New Entertainment Success")
	return newEvent
}

func UpdateEntertainent(id int, title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string) Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var updateEvent Entertainment
	database.
		Model(&updateEvent).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"Title":     title,
			"Price":     price,
			"Location":  location,
			"Latitude":  latitude,
			"Longitude": longitude,
			"Date":      date.Add(-7 * time.Hour),
			"Type":      category,
			"Image":     image,
		})

	log.Println("Update Entertainment Success")
	return updateEvent
}

func DeleteEntertainment(id int) *Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var event Entertainment
	event = GetEntertainmentByID(id)

	if event.ID == 0 {
		log.Println("Delete Blog Failed")
		return &event
	}

	err := database.Delete(event).Error

	if err != nil {
		panic("Error Delete Entertainment !" + err.Error())
	}

	log.Println("Delete Entertainment Success")
	return &event
}
