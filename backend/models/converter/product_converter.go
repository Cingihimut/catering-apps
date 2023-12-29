package converter

type UserCartProduct struct {
	ID          uint           `json:"id"`
	ProductName string         `json:"product_name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Quantity    uint           `json:"quantity"`
	Images      []ProductImage `json:"images" gorm:"-"`
}

type ProductImage struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}
