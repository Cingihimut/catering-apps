package repositories

import (
	"github.com/Cingihimut/catering-apps/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.Users) (*models.Users, error) {
	result := r.DB.Exec("INSERT INTO users (email, user_name, password) VALUES (?, ?, ?)", user.Email, user.UserName, user.Password)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.Users, error) {
	var user models.Users

	result := r.DB.Raw("SELECT id, email, user_name, password FROM users WHERE email = ?", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
