package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Driver struct {
	gorm.Model
	Uuid             string `gorm:"unique_index;"`
	Phone            string `gorm:"unique; not null"`
	Password         string `gorm:"not null"`
	FirstName        string `gorm:"not null"`
	LastName         string `gorm:"not null"`
	Address          string
	DriverLicenseID  string
	DriverLicenseImg string
	ProfileImg       string
	IsVerify         bool `gorm:"default:true; not null"`
	IsActive         bool `gorm:"default:true; not null"`
}

func (d *Driver) GenerateEncryptedPassword() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(d.Password), 14)
	return string(hash)
}
