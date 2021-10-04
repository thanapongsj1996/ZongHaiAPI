package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
	"zonghai-api/config"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1632663073CreateCustomerRequestTable(),
			m1633342866CreateDriverTable(),
			m1633374146CreateDriverRequestTable(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
