package models

import (
	"time"
)

type KosContact struct {
	KosID       string    `gorm:"column:kos_id;type:uuid;primaryKey" json:"kos_id"`
	ContactName string    `gorm:"column:contact_name" json:"contact_name"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Whatsapp    *string   `gorm:"column:whatsapp" json:"whatsapp,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Kos Kos `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
}

func (KosContact) TableName() string {
	return "kos_contacts"
}
