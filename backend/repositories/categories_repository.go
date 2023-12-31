package repositories

import (
	"time"

	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (c *CategoryRepository) Create(category *models.Categories) (*models.Categories, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	query := "INSERT INTO categories (name, created_at, updated_at) VALUES (?,?,?) "
	result := c.DB.Exec(query, category.Name, category.CreatedAt, category.UpdatedAt)

	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (c *CategoryRepository) Update(updatedCategoryInput *models.Categories) (*models.Categories, error) {
	query := "UPDATE categories SET name = ? WHERE id = ? RETURNING id"
	result := c.DB.Exec(query, updatedCategoryInput.Name, updatedCategoryInput.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return updatedCategoryInput, nil
}

func (c *CategoryRepository) Delete(id uint) error {
	query := "DELETE FROM categories WHERE id = ?"
	result := c.DB.Exec(query, id)
	return result.Error
}

func (c *CategoryRepository) GetById(id uint) (*models.Categories, error) {
	var category models.Categories

	query := "SELECT * FROM categories WHERE id = ?"
	result := c.DB.Raw(query, id).Scan(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (c *CategoryRepository) GetAll() ([]models.Categories, error) {
	var categories []models.Categories

	query := "SELECT * FROM categories"
	result := c.DB.Raw(query).Scan(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
