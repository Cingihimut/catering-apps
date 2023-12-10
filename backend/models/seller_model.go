package models

import (
	"time"

	"gorm.io/gorm"
)

type Sellers struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email       string     `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	UserName    string     `gorm:"type:varchar(255)" json:"username"`
	Password    string     `gorm:"type:varchar(255)" json:"password"`
	UserPicture string     `gorm:"type:varchar(255)" json:"seller_picture"`
	UserBanner  string     `gorm:"type:varchar(255)" json:"seller_banner"`
	Address     Address    `gorm:"foreignKey:UserID" json:"address"`
	Caterings   []Caterings `gorm:"foreignKey:SellerID" json:"caterings"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}