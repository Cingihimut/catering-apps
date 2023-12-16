package models

import (
	"time"

	"gorm.io/gorm"
)

type EnumRole string

const (
	RoleUser  EnumRole = "User"
	RoleOwner EnumRole = "Owner"
)

type Users struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email       string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	Role        EnumRole       `gorm:"not null" json:"role"`
	Picture     string         `gorm:"type:varchar(255)" json:"picture"`
	PhoneNumber string         `gorm:"type:varchar(255)" json:"phone_number"`
	Address     []Address      `gorm:"foreignKey:UserID" json:"adresses"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" `
}
