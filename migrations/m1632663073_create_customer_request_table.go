package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1632663073CreateCustomerRequestTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1632663073",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.CustomerRequest{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("customer_requests")
		},
	}
}