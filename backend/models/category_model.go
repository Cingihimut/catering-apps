package models

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Products  []Products     `gorm:"many2many:product_categories" json:"-"`
}
