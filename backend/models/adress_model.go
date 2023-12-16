package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	IsMain      bool           `json:"is_main"`
	UserName    string         `gorm:"type:varchar(255)" json:"user_name"`
	Country     string         `gorm:"not null" json:"country"`
	Province    string         `gorm:"not null" json:"province"`
	City        string         `gorm:"not null" json:"city"`
	PostalCode  string         `gorm:"not null" json:"postal_code"`
	FullAddress string         `gorm:"not null" json:"full_address"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
