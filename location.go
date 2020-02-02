package main

import "time"

type Location struct {
	ID        int     `gorm:"PRIMARY_KEY"`
	Longitude float32 `gorm:"DECIMAL"`
	Latitude  float32 `gorm:"DECIMAL"`
	City      string  `gorm:"VARCHAR(100)"`
	Province  string  `gorm:"VARCHAR(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
