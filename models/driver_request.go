package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DriverRequest struct {
	gorm.Model
	Uuid             string    `gorm:"unique_index;"`
	Description      string    `gorm:"not null"`
	StartPrice       float64   `gorm:"not null"`
	Phone            string    `gorm:"not null"`
	DeparturePlace   string    `gorm:"not null"`
	DepartureTime    time.Time `gorm:"not null"`
	DestinationPlace string    `gorm:"not null"`
	DestinationTime  time.Time `gorm:"not null"`
	PlaceOnTheWay    string
	IsActive         bool `gorm:"default: false; not null"`
	DriverId         uint
	Driver           Driver
}
