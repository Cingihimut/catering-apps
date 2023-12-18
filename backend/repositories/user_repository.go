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
	query := "INSERT INTO users (email, name, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	result := r.DB.Exec(query, user.Email, user.Name, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(user *models.Users) (*models.Users, error) {
	resultUser := &models.Users{}
	query := "SELECT id, email, name, password, role FROM users WHERE email = ?"
	result := r.DB.Raw(query, user.Email).Scan(resultUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return resultUser, nil
}

func (r *UserRepository) FindById(id uint) (*models.Users, error) {
	var user models.Users

	query := "SELECT * password FROM users WHERE id = ?"
	result := r.DB.Raw(query, id).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
