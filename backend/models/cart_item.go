package models


type CartItem struct {
	ProductID uint `gorm:"primaryKey" json:"product_id"`
	CartID    uint `gorm:"primaryKey" json:"cart_id"`
	Quantity  uint `json:"quantity"`
}
