package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044814CreateKosContactsTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029044814CreateKosContactsTable) Signature() string {
	return "20261029044814_create_kos_contacts_table"
}

// Up Run the migrations.
func (r *M20261029044814CreateKosContactsTable) Up() error {
	if !facades.Schema().HasTable("kos_contacts") {
		return facades.Schema().Create("kos_contacts", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Uuid("kos_id")
			table.String("contact_name")
			table.String("phone")
			table.String("whatsapp").Nullable()
			table.Primary("id")
			table.Index("kos_id")
			table.Foreign("kos_id").References("id").On("kos")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029044814CreateKosContactsTable) Down() error {
	return facades.Schema().DropIfExists("kos_contacts")
}
