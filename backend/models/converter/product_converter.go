// models/converter/all_products_response.go

package converter

import (
	"time"

	"github.com/Cingihimut/catering-apps/models"
)

type AllProductsResponse struct {
	ID          uint       `json:"id"`
	ProductName string     `json:"product_name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Images      []Image    `json:"images" gorm:"-"`
	Categories  []Category `json:"categories" gorm:"-"`
}

type Image struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ConvertToAllProductsResponse(product *models.Products) *AllProductsResponse {
	var images []Image
	for _, image := range product.Images {
		images = append(images, Image{
			ID:        image.ID,
			ProductID: image.ProductID,
			ImageURL:  image.ImageURL,
		})
	}

	var categories []Category
	for _, category := range product.Categories {
		categories = append(categories, Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return &AllProductsResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Images:      images,
		Categories:  categories,
	}
}
