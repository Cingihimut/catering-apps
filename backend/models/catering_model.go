package models

import (
	"time"

	"gorm.io/gorm"
)

type Caterings struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	SellerID     uint           `json:"seller_id"`
	Seller       Sellers        `gorm:"foreignKey:SellerID;constraint:OnDelete:CASCADE" json:"seller"`
	CateringName string         `json:"catering_name"`
	Description  string         `json:"description"`
	Price        float64        `json:"price"`
	Images       []CateringImages `gorm:"foreignKey:CateringID" json:"images"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

