package user

import (
	"../../connection"
	"log"
	"time"
)

type Admin struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Email     string `gorm:"VARCHAR(100); NOT NULL"`
	Password  string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Admin{})

	log.Println("Initialize Admin Success")
}

func DropTableAdmin() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&Admin{})
	database.AutoMigrate(&Admin{})

	log.Println("Drop Database Success")
}

func GetAllAdmin() []Admin {
	database := connection.GetConnection()
	defer database.Close()

	var admins []Admin
	database.Find(&admins)

	return admins
}

func InsertAdmin(email string, password string) *Admin {
	database := connection.GetConnection()
	defer database.Close()

	newAdmin := &Admin{
		Email:    email,
		Password: password,
	}
	database.Save(newAdmin)

	log.Println("Insert New Admin Success")
	return newAdmin
}

func ValidateAdminLogin(email string, password string) Admin {
	database := connection.GetConnection()
	defer database.Close()

	var admin Admin
	database.
		Where("email = ? and password = ?", email, password).
		First(&admin)

	return admin
}
