package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1633374146CreateDriverRequestTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1633374146",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.DriverJob{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("driver_requests")
		},
	}
}
