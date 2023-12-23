package models

type ProductImages struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint   `gorm:"index" json:"product_id"`
	ImageURL  string `gorm:"type:varchar(255);not null" json:"image_url"`
}
