package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DriverJob struct {
	gorm.Model
	Uuid                       string    `gorm:"unique_index;"`
	Description                string    `gorm:"not null"`
	StartPrice                 float64   `gorm:"not null"`
	Phone                      string    `gorm:"not null"`
	DeparturePlace             string    `gorm:"not null"`
	DepartureTime              time.Time `gorm:"not null" time_format:"2006-01-02 3:04PM"`
	DestinationPlace           string    `gorm:"not null"`
	DestinationTime            time.Time `gorm:"not null" time_format:"2006-01-02 3:04PM"`
	PlaceOnTheWay              string
	IsActive                   bool `gorm:"default: true; not null"`
	DriverId                   uint
	Driver                     Driver
	DriverJobDeliveryResponses []DriverJobDeliveryResponse
}
