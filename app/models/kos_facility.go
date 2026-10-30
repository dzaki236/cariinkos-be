package models

import "github.com/goravel/framework/database/orm"

type KosFacility struct {
	orm.Model
	KosID      string `gorm:"column:kos_id;type:uuid;primaryKey" json:"kos_id"`
	FacilityID string `gorm:"column:facility_id;type:uuid;primaryKey" json:"facility_id"`
}

func (KosFacility) TableName() string {
	return "kos_facilities"
}
