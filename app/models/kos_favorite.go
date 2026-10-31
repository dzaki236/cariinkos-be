package models

import "time"

type KosFavorite struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID    string    `gorm:"column:user_id;type:uuid;uniqueIndex:idx_user_kos_favorite" json:"user_id"`
	KosID     string    `gorm:"column:kos_id;type:uuid;uniqueIndex:idx_user_kos_favorite" json:"kos_id"`
	Note      string    `gorm:"column:note" json:"note"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Kos       Kos       `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
}

func (KosFavorite) TableName() string { return "kos_favorites" }
