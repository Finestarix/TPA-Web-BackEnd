package models

import "time"

type Airport struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Code      string `gorm:"CHAR(3); NOT NULL"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	City      string `gorm:"VARCHAR(100); NOT NULL"`
	CityCode  string `gorm:"VARCHAR(4); NOT NULL"`
	Province  string `gorm:"VARCHAR(100); NOT NULL"`
	Country   string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
