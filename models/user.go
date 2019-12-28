package models

import (
	"../connection"
	"log"
	"time"
)

type User struct {
	ID        int 		`gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"VARCHAR(100); NOT NULL"`
	Email     string     `gorm:"VARCHAR(100); NOT NULL"`
	Phone     string     `gorm:"VARCHAR(100); NOT NULL"`
	Password  string     `gorm:"VARCHAR(100); NOT NULL"`
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

func GetUser(id int) User {
	database := connection.GetConnection()
	defer database.Close()

	var user []User
	database.
		Where("id = ?", id).
		Find(&user)
	if len(user) == 0 {
		panic("There is No User")
	}

	return user[0]
}

func InsertUser(name string, email string, phone string, password string) (*User, error) {
	database := connection.GetConnection()
	defer database.Close()

	newUser := &User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	database.Save(newUser)

	log.Println("Insert New User Success")
	return newUser, nil
}

func UpdateUser(id int, name string, email string, phone string, password string) *User {
	database := connection.GetConnection()
	defer database.Close()

	var user User
	database.
		Model(&user).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":     name,
			"email":    email,
			"phone":    phone,
			"password": password,
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

	var user User = GetUser(id)
	error := database.Delete(user).Error
	if error != nil {
		panic("Error Delete User !" + error.Error())
	}

	log.Println("Delete User Success")
	return &user
}
