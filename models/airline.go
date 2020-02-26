package models

import "time"

type Airline struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	Icon      string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
