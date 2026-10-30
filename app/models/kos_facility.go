package models

import (
	"time"
)

type KosFacility struct {
	KosID      string    `gorm:"column:kos_id;type:uuid;primaryKey" json:"kos_id"`
	FacilityID string    `gorm:"column:facility_id;type:uuid;primaryKey" json:"facility_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	Kos        Kos       `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
	Facility   Facility  `gorm:"foreignKey:FacilityID;references:ID" json:"facility,omitempty"`
}

func (KosFacility) TableName() string {
	return "kos_facilities"
}
