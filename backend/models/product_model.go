package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID          uint            `gorm:"primaryKey;autoIncrement" json:"id" `
	ProductName string          `json:"product_name" form:"product_name"`
	Description string          `json:"description" form:"description"`
	Price       float64         `json:"price" form:"price"`
	Images      []ProductImages `gorm:"foreignKey:ProductID" json:"images"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	Categories  []Categories    `gorm:"many2many:product_categories" json:"categories"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}
