package models

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID                uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string              `gorm:"type:varchar(255)" json:"name"`
	ProductCategories []ProductCategories `gorm:"foreignKey:CategoryID" json:"-"`
	CreatedAt         time.Time           `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt         time.Time           `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt      `gorm:"index" json:"deleted_at,omitempty"`
}
