package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044824CreateKosPhotosTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029044824CreateKosPhotosTable) Signature() string {
	return "20261029044824_create_kos_photos_table"
}

// Up Run the migrations.
func (r *M20261029044824CreateKosPhotosTable) Up() error {
	if !facades.Schema().HasTable("kos_photos") {
		return facades.Schema().Create("kos_photos", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Uuid("kos_id")
			table.String("image_url")
			table.String("public_id").Nullable()
			table.Primary("id")
			table.Index("kos_id")
			table.Foreign("kos_id").References("id").On("kos")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029044824CreateKosPhotosTable) Down() error {
	return facades.Schema().DropIfExists("kos_photos")
}
