package converter

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type ImageSlice []struct {
	ImageURL string `json:"image_url"`
}

func (is ImageSlice) Value() (driver.Value, error) {
	return json.Marshal(is)
}

func (is *ImageSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), is)
}

type CategorySlice []struct {
	Name string `json:"name"`
}

func (cs CategorySlice) Value() (driver.Value, error) {
	return json.Marshal(cs)
}

func (cs *CategorySlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), cs)
}

type UserCartProduct struct {
	ID          uint          `json:"id"`
	ProductName string        `json:"product_name"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Images      ImageSlice    `gorm:"type:jsonb" json:"images"`
	Categories  CategorySlice `gorm:"type:jsonb" json:"categories"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
