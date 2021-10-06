package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1633545059CreateDriverJobPreOrderTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1633545059",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.DriverJobPreOrder{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("driver_job_pre_orders")
		},
	}
}
