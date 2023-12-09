package models

import (
	"time"

	"gorm.io/gorm"
)

type Sellers struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email       string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	UserName    string `gorm:"type:varchar(255)" json:"username"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	UserPicture string `gorm:"type:varchar(255)" json:"seller_picture"`
	UserBanner  string `gorm:"type:varchar(255)" json:"seller_banner"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
