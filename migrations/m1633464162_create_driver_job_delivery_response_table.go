package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1633464162CreateDriverJobDeliveryResponseTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1633464162",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.DriverJobDeliveryResponse{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("driver_job_delivery_responses")
		},
	}
}
