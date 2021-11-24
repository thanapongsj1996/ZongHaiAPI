package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
	"zonghai-api/config"
)

/*
	date +%s to get timestamp from terminal
*/

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1632663073CreateCustomerJobTable(),
			m1633342866CreateDriverTable(),
			m1633374146CreateDriverJobDeliveryTable(),
			m1633464162CreateDriverJobDeliveryResponseTable(),
			m1633545059CreateDriverJobPreOrderTable(),
			m1633545103CreateDriverJobPreOrderResponseTable(),
			m1637768912CreateProvidedJobsTable(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
