package repositories

import (
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
	SELECT
    p.id ,
    p.product_name,
    p.description,
    p.price,
    ci.quantity,
    JSON_AGG(DISTINCT json_build_object(
        'id', i.id,
        'product_id', i.product_id,
        'image_url', i.image_url
    )::jsonb) AS images
FROM
    carts c
JOIN
    cart_items ci ON c.id = ci.cart_id
JOIN
    products p ON ci.product_id = p.id
LEFT JOIN
    product_images i ON p.id = i.product_id
WHERE
    c.user_id = ?
GROUP BY
    p.id, product_name, description, price, ci.quantity;

	`

	if err := r.DB.Raw(rawSQL, userID).Scan(&userCartProducts).Error; err != nil {
		return nil, err
	}

	return userCartProducts, nil

}
