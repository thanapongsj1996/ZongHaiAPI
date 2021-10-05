package models

import "github.com/jinzhu/gorm"

type CustomerJob struct {
	gorm.Model
	Title           string  `gorm:"not null"`
	Description     string  `gorm:"not null"`
	Price           float64 `gorm:"not null"`
	Phone           string  `gorm:"not null"`
	PickupLocation  string  `gorm:"not null"`
	DeliverLocation string  `gorm:"not null"`
}
