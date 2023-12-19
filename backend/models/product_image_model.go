package models

type ProductImages struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"url"`
}
