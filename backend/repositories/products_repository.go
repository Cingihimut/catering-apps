package repositories

import (
	"fmt"
	"log"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/models/converter"
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
	query := "INSERT INTO products (product_name, description, price) VALUES (?, ?, ?) RETURNING id"
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

func (c *ProductRepository) SaveCategories(tx *gorm.DB, productID uint, categories []models.Categories) error {
	tx.Exec("DELETE FROM product_categories WHERE product_id = ?", productID)

	for _, category := range categories {
		tx.Exec("INSERT INTO product_categories (product_id, category_id) VALUES (?, ?)", productID, category.ID)
	}

	return nil
}

func (c *ProductRepository) SaveImage(tx *gorm.DB, productImages *models.ProductImages) (*models.ProductImages, error) {
	query := "INSERT INTO product_images (product_id, image_url) VALUES (?, ?) RETURNING id"
	result := tx.Raw(query, productImages.ProductID, productImages.ImageURL).Row()

	if result.Err() != nil {
		return nil, result.Err()
	}

	var insertedID uint
	if err := result.Scan(&insertedID); err != nil {
		return nil, err
	}
	productImages.ID = insertedID

	return productImages, nil
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
func (r *ProductRepository) GetAllProducts() ([]converter.AllProductsResponse, error) {
	query := `
        SELECT
            p.*,
            COALESCE(STRING_AGG(DISTINCT c.name, ', '), '') AS categories,
            COALESCE(STRING_AGG(DISTINCT pi.image_url, ', '), '') AS images
        FROM
            products p
        LEFT JOIN
            product_categories pc ON p.id = pc.product_id
        LEFT JOIN
            categories c ON pc.category_id = c.id
        LEFT JOIN
            product_images pi ON p.id = pi.product_id
        GROUP BY
            p.id, p.product_name;
    `

	var productsWithDetails []converter.AllProductsResponse
	rows, err := r.DB.Raw(query).Rows()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productDetails converter.AllProductsResponse
		err := r.DB.ScanRows(rows, &productDetails)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		productsWithDetails = append(productsWithDetails, productDetails)
	}

	return productsWithDetails, nil
}
