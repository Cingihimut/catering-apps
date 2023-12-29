package models

type CartItems struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint `gorm:"index" json:"product_id"`
	OwnerID   uint `gorm:"index" json:"owner_id"`
	Quantity  uint `gorm:"not null" json:"quantity"`
}
