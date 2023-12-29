package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID          uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName string          `json:"product_name" validate:"required,max=255"`
	Description string          `json:"description" form:"description"`
	Price       float64         `json:"price"`
	Images      []ProductImages `gorm:"foreignKey:ProductID" json:"images"`
	Categories  []Categories    `gorm:"many2many:product_categories;" json:"categories"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}
