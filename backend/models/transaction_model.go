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
	ID               uint               `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderDate        time.Time          `gorm:"type:date;not null" json:"order_date"`
	OrderTime        time.Time          `gorm:"type:time;not null" json:"order_time"`
	AddressID        uint               `gorm:"index" json:"address_id"`
	Total            float64            `gorm:"not null" json:"total"`
	PaymentMethod    string             `gorm:"default:null" json:"payment_method"`
	Delivery         bool               `gorm:"default:false" json:"delivery"`
	Status           EnumStatus         `gorm:"default:'pending'" json:"status"`
	TransactionItems []TransactionItems `gorm:"foreignKey:TransactionID" json:"-"`
	CreatedAt        time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"deleted_at,omitempty"`
}
