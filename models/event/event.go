package event

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type Entertainment struct {
	ID            uint `gorm:"primary_key"`
	Title         string
	Price         int
	Location      string
	Latitude      float64
	Longitude     float64
	Date          time.Time
	Type          string
	Image         string
	Description   string
	TermCondition string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
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

func GetEntertainmentByCategory(category string) []Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var event []Entertainment
	database.Where("type = ?", category).Limit(3).Find(&event)

	return event
}

func GetFilterEntertainment(lowPrice int, highPrice int,
	startDate time.Time, endDate time.Time, isActivity string,
	isAttraction string, isEvent string, location string) []Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	eventCat := ""
	if isEvent == "true" {
		eventCat = "Event"
	}

	attractionCat := ""
	if isAttraction == "true" {
		attractionCat = "Attraction"
	}

	activityCat := ""
	if isActivity == "true" {
		activityCat = "Activity"
	}

	locationStr := "%" + location + "%"

	var event []Entertainment
	if core.ValidateAPIKey() == false {
		return event
	}
	database.Where("price <= ? AND price >= ? AND DATE_PART('day', date) >= ? AND " +
		"DATE_PART('day', date) <= ? AND type IN (?) AND location LIKE ?",
		highPrice, lowPrice, startDate.Day(), endDate.Day(),
		[]string{eventCat, attractionCat, activityCat}, locationStr).Find(&event)

	return event
}

func InsertEntertainment(title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string, description string, termCondition string) *Entertainment {

	database := connection.GetConnection()
	defer database.Close()

	if price <= 0 {
		return nil
	}

	newEvent := &Entertainment{
		Title:         title,
		Price:         price,
		Location:      location,
		Latitude:      latitude,
		Longitude:     longitude,
		Date:          date.Add(-7 * time.Hour),
		Type:          category,
		Image:         image,
		TermCondition: termCondition,
		Description:   description,
	}
	database.Save(newEvent)

	log.Println("Insert New Entertainment Success")
	return newEvent
}

func UpdateEntertainment(id int, title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string, description string, termCondition string) Entertainment {
	database := connection.GetConnection()
	defer database.Close()

	var updateEvent Entertainment
	database.
		Model(&updateEvent).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"Title":         title,
			"Price":         price,
			"Location":      location,
			"Latitude":      latitude,
			"Longitude":     longitude,
			"Date":          date.Add(-7 * time.Hour),
			"Type":          category,
			"Image":         image,
			"Description":   description,
			"TermCondition": termCondition,
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
