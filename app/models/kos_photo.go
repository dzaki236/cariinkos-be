package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type KosPhoto struct {
	orm.Model
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	KosID     string    `gorm:"column:kos_id;type:uuid;index" json:"kos_id"`
	ImageURL  string    `gorm:"column:image_url" json:"image_url"`
	PublicID  *string   `gorm:"column:public_id" json:"public_id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Kos Kos `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
}

func (KosPhoto) TableName() string {
	return "kos_photos"
}
