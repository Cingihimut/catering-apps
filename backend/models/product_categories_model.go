package models

type ProductCategories struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}
