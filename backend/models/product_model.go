package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID                uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName       string              `gorm:"type:varchar(255);not null" json:"product_name" form:"product_name"`
	Description       string              `gorm:"type:varchar(255);" json:"description;" form:"description"`
	Price             float64             `gorm:"type:varchar(255);not null" json:"price" form:"price"`
	Images            []ProductImages     `gorm:"foreignKey:ProductID" json:"images"`
	ProductCategories []ProductCategories `gorm:"foreignKey:ProductID" json:"-"`
	CartItems         []CartItems         `gorm:"foreignKey:ProductID" json:"-"`
	TransactionItems  []TransactionItems  `gorm:"foreignKey:TransactionID" json:"-"`
	CreatedAt         time.Time           `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt         time.Time           `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt      `gorm:"index" json:"deleted_at,omitempty"`
}
