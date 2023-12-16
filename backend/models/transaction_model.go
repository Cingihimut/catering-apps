package models

import (
	"time"

	"gorm.io/gorm"
)

type EnumStatus string

const (
	StatusPending    EnumStatus = "Pending"
	StatusProcessing EnumStatus = "Processing"
	StatusCompleted  EnumStatus = "Completed"
	StatusCancelled  EnumStatus = "Cancelled"
)

type Transactions struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderDate     time.Time      `json:"order_date"`
	OrderTime     time.Time      `json:"order_time"`
	BuyerID       uint           `gorm:"foreignKey:buyer_id;" json:"buyer_id"`
	AddressId     uint           `gorm:"foreignKey:address_id" json:"address_id"`
	CartID        uint           `gorm:"foreignKey:cart_id" json:"cart_id"`
	Total         float64        `json:"total"`
	PaymentMethod string         `json:"payment_method"`
	Status        EnumStatus     `json:"status"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
