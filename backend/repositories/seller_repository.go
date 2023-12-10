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
	query := "INSERT INTO sellers (email, user_name, password) VALUES (?, ?, ?)"
	result := r.DB.Exec(query, seller.Email, seller.UserName, seller.Password)
	if result.Error != nil {
		return nil, result.Error
	}

	return seller, nil
}

func (r *SellerRepository) FindByEmail(email string) (*models.Sellers, error) {
	var seller models.Sellers
	
	query := "SELECT id, email, user_name, password FROM sellers WHERE email = ?"
	result := r.DB.Raw(query, email).Scan(&seller)
	if result.Error != nil {
		return nil, result.Error
	}

	return &seller, nil
}

func (r *SellerRepository) FindById(id uint) (*models.Sellers, error){
	var seller models.Sellers

	query := "SELECT * password FROM sellers WHERE id = ?"
	result := r.DB.Raw(query, id).Scan(&seller)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &seller, nil
}

