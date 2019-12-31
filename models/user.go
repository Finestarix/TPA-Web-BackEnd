package models

import (
	"../connection"
	"log"
	"time"
)

type User struct {
	ID             int    `gorm:"PRIMARY_KEY"`
	Title          string `gorm:"VARCHAR(100);"`
	FirstName      string `gorm:"VARCHAR(100); NOT NULL"`
	LastName       string `gorm:"VARCHAR(100); NOT NULL"`
	Email          string `gorm:"VARCHAR(100); NOT NULL; UNIQUE"`
	IsEmailConfirm int    `gorm:"INTEGER"`
	PhoneCode      string `gorm:"VARCHAR(100); NOT NULL"`
	Phone          string `gorm:"VARCHAR(100); NOT NULL"`
	IsPhoneConfirm int    `gorm:"INTEGER"`
	Password       string `gorm:"VARCHAR(100); NOT NULL"`
	City           string `gorm:"VARCHAR(100)"`
	Address        string `gorm:"VARCHAR(100)"`
	ZipCode        int    `gorm:"INTEGER"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.AutoMigrate(&User{})
	log.Println("Initialize User Success")
}

func DropTableUser() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&User{})
	database.AutoMigrate(&User{})
	log.Println("Drop Database Success")
}

func GetAllUser() []User {
	database := connection.GetConnection()
	defer database.Close()

	var users []User
	database.Find(&users)
	return users
}

func GetUserByID(id int) []User {
	database := connection.GetConnection()
	defer database.Close()

	var user []User
	database.
		Where("id = ?", id).
		Find(&user)

	return user
}

func GetUserByPhone(phone string) []User {
	database := connection.GetConnection()
	defer database.Close()

	var user []User
	database.
		Where("phone = ?", phone).
		Find(&user)

	log.Println("Someone Search User by Phone")

	return user
}

func GetUserByEmail(email string) []User {
	database := connection.GetConnection()
	defer database.Close()

	var user []User
	database.
		Where("email = ?", email).
		Find(&user)

	log.Println("Someone Search User by Email")

	return user
}

func InsertUser(firstName string, lastName string, email string, phoneCode string, phone string, password string) *User {
	database := connection.GetConnection()
	defer database.Close()

	newUser := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		PhoneCode: phoneCode,
		Phone:     phone,
		Password:  password,
	}
	database.Save(newUser)

	log.Println("Insert New User Success")
	return newUser
}

func UpdateUserProfile(id int, title string, firstName string, lastName string, city string, address string, zipCode int) *User {
	database := connection.GetConnection()
	defer database.Close()

	var user User
	database.
		Model(&user).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":     title,
			"first_name": firstName,
			"last_name":  lastName,
			"city":      city,
			"address":   address,
			"zip_code":   zipCode,
		})
	database.
		Where("id = ?", id).
		Find(&user)

	log.Println("Update User Success")
	return &user
}

func DeleteUser(id int) *User {
	database := connection.GetConnection()
	defer database.Close()

	var users []User = GetUserByID(id)

	if len(users) != 0 {
		return nil
	}

	var user = users[0]
	error := database.Delete(user).Error
	if error != nil {
		panic("Error Delete User !" + error.Error())
	}

	log.Println("Delete User Success")
	return &user
}
