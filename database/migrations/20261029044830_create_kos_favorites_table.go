package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044830CreateKosFavoritesTable struct{}

func (r *M20261029044830CreateKosFavoritesTable) Signature() string {
	return "20261029044830_create_kos_favorites_table"
}

func (r *M20261029044830CreateKosFavoritesTable) Up() error {
	if !facades.Schema().HasTable("kos_favorites") {
		return facades.Schema().Create("kos_favorites", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Uuid("user_id")
			table.Uuid("kos_id")
			table.String("note").Nullable()
			table.Primary("id")
			table.Unique("user_id", "kos_id")
			table.Foreign("user_id").References("id").On("users")
			table.Foreign("kos_id").References("id").On("kos")
			table.TimestampsTz()
		})
	}

	return nil
}

func (r *M20261029044830CreateKosFavoritesTable) Down() error {
	return facades.Schema().DropIfExists("kos_favorites")
}
