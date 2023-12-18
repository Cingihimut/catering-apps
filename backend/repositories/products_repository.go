package repositories

import (
	"fmt"
	"log"

	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (c *ProductRepository) Create(tx *gorm.DB, product *models.Products) (*models.Products, error) {
	query := "INSERT INTO products (product_name, description, price) VALUES (?, ?, ?, ?) RETURNING id"

	result := tx.Raw(query, product.ProductName, product.Description, product.Price).Row()

	if result.Err() != nil {
		return nil, result.Err()
	}

	var insertedID uint
	if err := result.Scan(&insertedID); err != nil {
		return nil, err
	}
	product.ID = insertedID
	return product, nil
}

func (c *ProductRepository) SaveImage(tx *gorm.DB, productImages *models.ProductImages) (*models.ProductImages, error) {
	query := "INSERT INTO product_images (product_id, image_url) VALUES (?,?) RETURNING id"
	result := tx.Raw(query, productImages.ProductID, productImages.ImageURL).Scan(&productImages.ID)

	if result.Error != nil {
		log.Printf("Error executing query: %v\n", result.Error)
		return nil, result.Error
	}
	return productImages, nil
}

func (c *ProductRepository) GetAllProductFromDB() ([]models.Products, error) {
	var products []models.Products
	result := c.DB.Raw("SELECT * FROM products").Scan(&products)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (c *ProductRepository) LoadImages(product *models.Products) {
	query := "SELECT * FROM product_images WHERE product_id = ?"
	result := c.DB.Raw(query, product.ID).Scan(&product.Images)

	if result.Error != nil {
		log.Printf("Error loading images: %v\n", result.Error)
	}
}

func (r *ProductRepository) GetProductByID(productID uint) (*models.Products, error) {
	var product models.Products

	result := r.DB.Raw("SELECT * FROM products WHERE seller_id = ?", productID).Scan(&product)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}