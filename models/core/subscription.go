package core

import (
	"../../email"
	"../../connection"
	"log"
	"time"
)

type Subscription struct {
	ID int  `gorm:"PRIMARY_KEY"`
	Email string `gorm:"VARCHAR(100)"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.AutoMigrate(&Subscription{})

	log.Println("Initialize Subscription Success")
}

func InsertNewSubscription(newEmail string) *Subscription {
	database := connection.GetConnection()
	defer database.Close()

	newSubscription := &Subscription{
		Email: newEmail,
	}
	database.Save(newSubscription)

	to := []string{newEmail}
	email.SubscriptionEmail(to)

	log.Println("Insert New Subscription Success")
	return newSubscription
}

func SendSubscriptionToAll() []Subscription {
	database := connection.GetConnection()
	defer database.Close()

	var subscription []Subscription
	database.Find(&subscription)

	for i, _ := range subscription {
		to := []string{subscription[i].Email}
		email.SubscriptionEmail(to)
	}

	return subscription
}