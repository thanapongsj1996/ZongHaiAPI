package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1633545103CreateDriverJobPreOrderResponseTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1633545103",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.DriverJobPreOrderResponse{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("driver_job_pre_order_responses")
		},
	}
}
