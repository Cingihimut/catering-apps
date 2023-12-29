package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"cart_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CartItems []CartItem     `gorm:"not null" json:"cart_items"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
