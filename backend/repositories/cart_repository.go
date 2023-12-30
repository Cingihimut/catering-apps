package repositories

import (
	"encoding/json"
	"log"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/models/converter"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}
func (r *CartRepository) AddToCart(userID, productID, quantity uint) error {
	var cartID uint
	query := "SELECT id FROM carts WHERE user_id = ?"
	result := r.DB.Raw(query, userID).Row()
	if err := result.Scan(&cartID); err != nil {
		if err := r.DB.Create(&models.Cart{UserID: userID}).Error; err != nil {
			return err
		}
		query = "SELECT id FROM carts WHERE user_id = ?"
		result = r.DB.Raw(query, userID).Row()
		if err := result.Scan(&cartID); err != nil {
			return err
		}
	}

	query = "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?) " +
		"ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity"
	return r.DB.Exec(query, cartID, productID, quantity).Error
}

func (r *CartRepository) GetUserCart(userID uint) ([]converter.UserCartProduct, error) {
	var userCartProducts []converter.UserCartProduct

	rawSQL := `
	SELECT DISTINCT ON (products.id) 
			products.id, 
			products.product_name, 
			products.description, 
			products.price, 
			JSON_AGG(DISTINCT json_build_object('image_url', product_images.image_url)::jsonb) AS Images, 
			JSON_AGG(DISTINCT json_build_object('name', categories.name)::jsonb) AS categories,
			products.created_at, 
			products.updated_at
		FROM products
		JOIN cart_items ON products.id = cart_items.product_id
		JOIN carts ON cart_items.cart_id = carts.id
		LEFT JOIN product_images ON products.id = product_images.product_id
		LEFT JOIN product_categories ON products.id = product_categories.product_id
		LEFT JOIN categories ON product_categories.category_id = categories.id
		WHERE carts.user_id = ?
		GROUP BY products.id, products.product_name, products.description, products.price, products.created_at, products.updated_at
		ORDER BY products.id;

	`

	if err := r.DB.Debug().Raw(rawSQL, userID).Scan(&userCartProducts).Error; err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(userCartProducts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("JSON DATA : %v", string(jsonData))
	return userCartProducts, nil

}
