package core

import (
	"../../connection"
	"../../middleware"
	"log"
	"time"
)

type API struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Key       string `gorm:"VARCHAR(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.AutoMigrate(&API{})

	log.Println("Initialize API Success")
}

func ValidateAPIKey() bool {

	database := connection.GetConnection()
	defer database.Close()

	var APIKey API
	database.Where("Key = ?", middleware.APIKey).First(&APIKey)

	if APIKey.ID == 0 {
		return false
	}

	return true
}
