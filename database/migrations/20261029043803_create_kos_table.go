package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029043803CreateKosTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029043803CreateKosTable) Signature() string {
	return "20261029043803_create_kos_table"
}

// Up Run the migrations.
func (r *M20261029043803CreateKosTable) Up() error {
	if !facades.Schema().HasTable("kos") {
		return facades.Schema().Create("kos", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Uuid("user_id")
			table.String("name")
			table.Text("description").Nullable()
			table.Text("address")
			table.String("province")               //provinsi
			table.String("city")                   //kota
			table.String("district")               //kecamatan
			table.String("subdistrict").Nullable() //kelurahan
			table.String("gmaps_link").Nullable()  //link gmaps
			table.UnsignedBigInteger("price_monthly")
			table.String("type").Default("mixed") //mixed, man, woman
			table.String("status").Default("draft")
			table.UnsignedBigInteger("room_size").Nullable()
			table.UnsignedBigInteger("total_rooms").Default(0)
			table.UnsignedBigInteger("available_rooms").Default(0)
			table.UnsignedBigInteger("minimum_rent_months").Default(1)
			table.Text("rules").Nullable()
			table.Primary("id")
			table.Index("user_id")
			table.Index("city")
			table.Index("price_monthly")
			table.Foreign("user_id").References("id").On("users")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029043803CreateKosTable) Down() error {
	return facades.Schema().DropIfExists("kos")
}
