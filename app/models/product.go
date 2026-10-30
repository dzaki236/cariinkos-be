package models

import (
	"github.com/goravel/framework/database/orm"
)

type Product struct {
	orm.Model
	ID          string  `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description *string `gorm:"column:description" json:"description,omitempty"`
	Price       int64   `gorm:"column:price" json:"price"`
}

func (Product) TableName() string {
	return "products"
}
