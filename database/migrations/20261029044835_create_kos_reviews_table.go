package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029044835CreateKosReviewsTable struct{}

func (r *M20261029044835CreateKosReviewsTable) Signature() string {
	return "20261029044835_create_kos_reviews_table"
}

func (r *M20261029044835CreateKosReviewsTable) Up() error {
	if !facades.Schema().HasTable("kos_reviews") {
		return facades.Schema().Create("kos_reviews", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Uuid("user_id")
			table.Uuid("kos_id")
			table.UnsignedTinyInteger("rating")
			table.Text("comment").Nullable()
			table.Primary("id")
			table.Unique("user_id", "kos_id")
			table.Foreign("user_id").References("id").On("users")
			table.Foreign("kos_id").References("id").On("kos")
			table.TimestampsTz()
		})
	}

	return nil
}

func (r *M20261029044835CreateKosReviewsTable) Down() error {
	return facades.Schema().DropIfExists("kos_reviews")
}
