package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ProductID      uint    `gorm:"primaryKey;autoIncrement" json:"ProductID"`
	ProductName    string  `gorm:"type:varchar(255)" json:"ProductName"`
	Price          float64 `gorm:"type:numeric" json:"Price"`
	ProductPicture string  `gorm:"type:varchar(255)" json:"ProductPicture"`
	OwnerID        uint    `gorm:"index" json:"OwnerID"`
	Owner          Users   `gorm:"foreignKey:OwnerID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
