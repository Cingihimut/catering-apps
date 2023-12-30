package services

import (
	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
)

type CategoryService struct {
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(category *models.Categories) (*models.Categories, error) {
	return s.CategoryRepository.Create(category)
}

func (s *CategoryService) Update(updatedCategoryInput *models.Categories) (*models.Categories, error) {
	return s.CategoryRepository.Update(updatedCategoryInput)
}

func (s *CategoryService) Delete(id uint) error {
	return s.CategoryRepository.Delete(id)
}

func (s *CategoryService) GetById(id uint) (*models.Categories, error) {
	return s.CategoryRepository.GetById(id)
}

func (s *CategoryService) GetAll() ([]models.Categories, error) {
	return s.CategoryRepository.GetAll()
}
