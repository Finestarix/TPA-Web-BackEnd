package models

import "time"

type Route struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	From      string `gorm:"VARCHAR(100); NOT NULL"`
	To        string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
