package models

import "time"

type KosReview struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID    string    `gorm:"column:user_id;type:uuid;uniqueIndex:idx_user_kos_review" json:"user_id"`
	KosID     string    `gorm:"column:kos_id;type:uuid;uniqueIndex:idx_user_kos_review" json:"kos_id"`
	Rating    uint8     `gorm:"column:rating" json:"rating"`
	Comment   string    `gorm:"column:comment" json:"comment"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Kos       Kos       `gorm:"foreignKey:KosID;references:ID" json:"kos,omitempty"`
}

func (KosReview) TableName() string { return "kos_reviews" }
