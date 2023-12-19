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

func NewProductRepository(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: DB,
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

func (r *ProductRepository) GetAllProducts() ([]models.Products, error) {
	var products []models.Products

	// Eksekusi SQL mentah menggunakan Raw SQL di Gorm
	rawSQL := `
		SELECT 
			p.id,
			p.product_name,
			p.description,
			p.price,
			p.created_at,
			p.updated_at,
			i.id as image_id,
			i.product_id,
			i.image_url as image_url
		FROM 
			products p
		LEFT JOIN 
			product_images i ON p.id = i.product_id
		ORDER BY 
			p.id, i.id
	`

	rows, err := r.DB.Raw(rawSQL).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Printf("Hasil : %v", rows)

	// Map untuk melacak produk yang telah diproses
	processedProducts := make(map[uint]*models.Products)

	for rows.Next() {
		var product models.Products
		var image models.ProductImages

		err := rows.Scan(
			&product.ID, &product.ProductName, &product.Description, &product.Price,
			&product.CreatedAt, &product.UpdatedAt,
			&image.ID, &image.ProductID, &image.ImageURL,
		)

		if err != nil {
			return nil, err
		}

		// Lakukan pengecekan apakah produk sudah diproses atau belum
		if _, ok := processedProducts[product.ID]; !ok {
			// Jika belum, tambahkan produk ke slice dan tandai sebagai diproses
			products = append(products, product)
			processedProducts[product.ID] = &products[len(products)-1]
		}

		// Tambahkan gambar ke produk yang sesuai
		if image.ID != 0 {
			processedProducts[product.ID].Images = append(processedProducts[product.ID].Images, image)
		}
	}

	// Ambil kategori untuk setiap produk
	for i := range products {
		categories, err := r.getCategoriesForProduct(products[i].ID)
		if err != nil {
			return nil, err
		}
		products[i].Categories = categories
	}

	return products, nil
}
func (r *ProductRepository) getCategoriesForProduct(productID uint) ([]models.Categories, error) {
	query := fmt.Sprintf(`
		SELECT c.id, c.name
		FROM categories c
		LEFT JOIN product_categories pc ON c.id = pc.category_id AND pc.product_id = %d
		WHERE pc.product_id IS NOT NULL OR pc.product_id IS NULL
	`, productID)

	rows, err := r.DB.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Categories
	for rows.Next() {
		var category models.Categories
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
