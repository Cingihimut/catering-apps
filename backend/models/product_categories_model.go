package models

type ProductCategories struct {
	ID         uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID uint `gorm:"index" json:"category_id"`
	ProductID  uint `gorm:"index" json:"product_id"`
}
