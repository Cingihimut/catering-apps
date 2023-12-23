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
	Password    string         `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Role        EnumRole       `gorm:"not null;default:'user'" json:"role,omitempty"`
	Picture     string         `gorm:"type:varchar(255)" json:"picture,omitempty"`
	PhoneNumber string         `gorm:"type:varchar(255)" json:"phone_number,omitempty"`
	Addresses   []Addresses    `gorm:"foreignKey:OwnerID" json:"adresses"`
	CartItems   []CartItems    `gorm:"foreignKey:OwnerID" json:"-"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
