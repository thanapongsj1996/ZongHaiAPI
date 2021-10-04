package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1633342866CreateDriverTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1633342866",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Driver{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("drivers")
		},
	}
}
