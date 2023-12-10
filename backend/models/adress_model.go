package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`  
	UserName    string         `gorm:"type:varchar(255)" json:"user_name"`
	PhoneNumber string         `json:"phone_number"`
	Country     string         `gorm:"not null" json:"country"`
	Province    string         `gorm:"not null" json:"province"`
	City        string         `gorm:"not null" json:"city"`
	PostalCode  string         `gorm:"not null" json:"postal_code"`
	FullAddress string         `gorm:"not null" json:"full_address"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}