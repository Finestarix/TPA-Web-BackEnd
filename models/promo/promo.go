package promo

import (
	"../../connection"
	"log"
	"time"
)

type Promo struct {
	ID        int `gorm:"PRIMARY_KEY"`
	Name      string
	Detail    string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Promo{})

	log.Println("Initialize Promo Success")
}

func GetPromoBYID(id int) Promo {
	database := connection.GetConnection()
	defer database.Close()

	var promo Promo
	database.Where("id = ?", id).First(&promo)

	return promo
}

func GetAnotherPromo(id int) []Promo {
	database := connection.GetConnection()
	defer database.Close()

	var promo []Promo
	database.Where("id != ?", id).Limit(4).Find(&promo)

	return promo
}
