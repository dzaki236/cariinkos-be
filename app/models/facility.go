package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Facility struct {
	orm.Model
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;uniqueIndex" json:"name"`
	Icon      *string   `gorm:"column:icon" json:"icon,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Kos []Kos `gorm:"many2many:kos_facilities;foreignKey:ID;joinForeignKey:FacilityID;references:ID;joinReferences:KosID" json:"kos,omitempty"`
}

func (Facility) TableName() string {
	return "facilities"
}
