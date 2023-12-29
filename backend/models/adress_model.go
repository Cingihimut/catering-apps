package models

import (
	"time"

	"gorm.io/gorm"
)

type Addresses struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint           `gorm:"index;not null" json:"user_id"`
	IsMain      bool           `gorm:"default:false" json:"is_main"`
	Country     string         `gorm:"type:varchar(255);not null" json:"country"`
	Province    string         `gorm:"type:varchar(255);not null" json:"province"`
	City        string         `gorm:"type:varchar(255);not null" json:"city"`
	PostalCode  string         `gorm:"type:varchar(255);not null" json:"postal_code"`
	FullAddress string         `gorm:"type:varchar(255);not null" json:"full_address"`
	Description string         `gorm:"default:null" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
