package models

import (
	"github.com/jinzhu/gorm"
)

type DriverJobDeliveryResponse struct {
	gorm.Model
	Uuid           string `gorm:"unique_index;"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Items          string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	PickupPlace    string `gorm:"not null"`
	DeliverPlace   string `gorm:"not null"`
	IsActive       bool   `gorm:"default: true; not null"`
	IsDriverAccept bool   `gorm:"default: false; not null"`
	DriverJobId    uint
	DriverJob      DriverJob
}
