package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DriverJobPreOrder struct {
	gorm.Model
	Uuid                       string    `gorm:"unique_index;"`
	Description                string    `gorm:"not null"`
	Price                      float64   `gorm:"not null"`
	Phone                      string    `gorm:"not null"`
	ShopPlace                  string    `gorm:"not null"`
	DepartureTime              time.Time `gorm:"not null"`
	DestinationPlace           string    `gorm:"not null"`
	DestinationTime            time.Time `gorm:"not null"`
	IsActive                   bool      `gorm:"default: true; not null"`
	DriverId                   uint
	Driver                     Driver
	DriverJobPreOrderResponses []DriverJobPreOrderResponse
}
