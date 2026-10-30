package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type KosContact struct {
	orm.Model
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	KosID     string    `gorm:"column:kos_id;type:uuid;index" json:"kos_id"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Whatsapp  *string   `gorm:"column:whatsapp" json:"whatsapp,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Kos Kos `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
}

func (KosContact) TableName() string {
	return "kos_contacts"
}
