package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Kos struct {
	orm.Model
	ID                string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID            string    `gorm:"column:user_id;type:uuid;index" json:"user_id"`
	Name              string    `gorm:"column:name" json:"name"`
	Description       *string   `gorm:"column:description" json:"description,omitempty"`
	Address           string    `gorm:"column:address" json:"address"`
	Province          string    `gorm:"column:province" json:"province"`
	City              string    `gorm:"column:city;index" json:"city"`
	District          string    `gorm:"column:district" json:"district"`
	Subdistrict       *string   `gorm:"column:subdistrict" json:"subdistrict,omitempty"`
	GmapsLink         *string   `gorm:"column:gmaps_link" json:"gmaps_link,omitempty"`
	PriceMonthly      uint64    `gorm:"column:price_monthly;index" json:"price_monthly"`
	RoomSize          *uint64   `gorm:"column:room_size" json:"room_size,omitempty"`
	TotalRooms        uint64    `gorm:"column:total_rooms" json:"total_rooms"`
	AvailableRooms    uint64    `gorm:"column:available_rooms" json:"available_rooms"`
	MinimumRentMonths uint64    `gorm:"column:minimum_rent_months" json:"minimum_rent_months"`
	Rules             *string   `gorm:"column:rules" json:"rules,omitempty"`
	Type              string    `gorm:"column:type" json:"type"`
	Status            string    `gorm:"column:status" json:"status"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`

	User       User         `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Facilities []Facility   `gorm:"many2many:kos_facilities;foreignKey:ID;joinForeignKey:KosID;references:ID;joinReferences:FacilityID" json:"facilities,omitempty"`
	Contacts   []KosContact `gorm:"foreignKey:KosID;references:ID" json:"contacts,omitempty"`
	Photos     []KosPhoto   `gorm:"foreignKey:KosID;references:ID" json:"photos,omitempty"`
}

func (Kos) TableName() string {
	return "kos"
}
