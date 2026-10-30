package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044809CreateKosFacilitiesTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029044809CreateKosFacilitiesTable) Signature() string {
	return "20261029044809_create_kos_facilities_table"
}

// Up Run the migrations.
func (r *M20261029044809CreateKosFacilitiesTable) Up() error {
	if !facades.Schema().HasTable("kos_facilities") {
		return facades.Schema().Create("kos_facilities", func(table schema.Blueprint) {
			table.Uuid("kos_id")
			table.Uuid("facility_id")
			table.Primary("kos_id", "facility_id")
			table.Foreign("kos_id").References("id").On("kos")
			table.Foreign("facility_id").References("id").On("facilities")
			table.TimestampsTz()

		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029044809CreateKosFacilitiesTable) Down() error {
	return facades.Schema().DropIfExists("kos_facilities")
}
