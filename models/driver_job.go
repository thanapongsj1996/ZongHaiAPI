package models

import (
	"github.com/jinzhu/gorm"
)

type ProvidedJob struct {
	gorm.Model
	Uuid             string `gorm:"unique_index;"`
	FirstName        string `gorm:"default: -; not null"`
	LastName         string `gorm:"default: -; not null"`
	Description      string
	Price            float64 `gorm:"default: 0; not null"`
	Phone            string  `gorm:"default: -; not null"`
	DeparturePlace   string  `gorm:"default: -; not null"`
	DestinationPlace string  `gorm:"default: -; not null"`
	PlaceOnTheWay    string
	IsActive         bool `gorm:"default: true; not null"`
}
