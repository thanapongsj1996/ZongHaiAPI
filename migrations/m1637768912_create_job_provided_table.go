package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"zonghai-api/models"
)

func m1637768912CreateProvidedJobsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1637768912",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.ProvidedJob{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("provided_jobs")
		},
	}
}
