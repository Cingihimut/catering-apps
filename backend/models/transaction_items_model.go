package models

type TransactionItems struct {
	ID            uint `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID uint `gorm:"index" json:"transaction_id"`
	ProductID     uint `json:"product_id"`
	Quantity      uint `json:"quantity"`
}
