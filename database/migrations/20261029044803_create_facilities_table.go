package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044803CreateFacilitiesTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029044803CreateFacilitiesTable) Signature() string {
	return "20261029044803_create_facilities_table"
}

// Up Run the migrations.
func (r *M20261029044803CreateFacilitiesTable) Up() error {
	if !facades.Schema().HasTable("facilities") {
		return facades.Schema().Create("facilities", func(table schema.Blueprint) {
			table.Uuid("id")
			table.String("name")
			table.String("icon").Nullable()
			table.Primary("id")
			table.Unique("name")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029044803CreateFacilitiesTable) Down() error {
	return facades.Schema().DropIfExists("facilities")
}
