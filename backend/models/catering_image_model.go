package models

import (
	"time"

	"gorm.io/gorm"
)

type CateringImages struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CateringID uint       `json:"-"`
	ImageURL   string     `json:"image_url"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}