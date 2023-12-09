package repositories

import (
	"github.com/Cingihimut/catering-apps/models"

	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		DB: db,
	}
}

func (r *SellerRepository) Create(seller *models.Sellers) (*models.Sellers, error) {
	result := r.DB.Exec("INSERT INTO sellers (email, user_name, password) VALUES (?, ?, ?)", seller.Email, seller.UserName, seller.Password)
	if result.Error != nil {
		return nil, result.Error
	}

	return seller, nil
}

func (r *SellerRepository) FindByEmail(email string) (*models.Sellers, error) {
	var seller models.Sellers

	result := r.DB.Raw("SELECT id, email, user_name, password FROM sellers WHERE email = ?", email).Scan(&seller)
	if result.Error != nil {
		return nil, result.Error
	}

	return &seller, nil
}
