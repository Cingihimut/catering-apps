package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserEmail     string `gorm:"type:varchar(100);uniqueIndex" json:"userEmail"`
	UserName      string `gorm:"type:varchar(255)" json:"userName"`
	UserPassword  string `gorm:"type:varchar(255)" json:"userPassword"`
	UserPrivilege string `gorm:"type:varchar(50)" json:"userPrivilege"`
	UserPicture   string `gorm:"type:varchar(255)" json:"userPicture"`
	UserBanner    string `gorm:"type:varchar(255)" json:"userBanner"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
