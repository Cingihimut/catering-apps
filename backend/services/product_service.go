package services

import (
	"mime/multipart"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/Cingihimut/catering-apps/utils"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(ProductRepository repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: ProductRepository,
	}
}
func (p *ProductService) Create(product *models.Products, imageURLs []string) (*models.Products, error) {
	tx := p.ProductRepository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	createdProduct, err := p.ProductRepository.Create(tx, product)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := p.ProductRepository.SaveCategories(tx, createdProduct.ID, product.Categories); err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, imageURL := range imageURLs {
		productImage := models.ProductImages{
			ProductID: createdProduct.ID,
			ImageURL:  imageURL,
		}

		if _, err := p.ProductRepository.SaveImage(tx, &productImage); err != nil {
			tx.Rollback()
			return nil, err
		}
		createdProduct.Images = append(createdProduct.Images, productImage)
	}

	tx.Commit()

	return createdProduct, nil
}

func (s *ProductService) SaveImages(files []*multipart.FileHeader) ([]string, error) {
	var imageURLs []string

	for _, file := range files {
		imageURL, err := utils.SaveImageToLocal(file)
		if err != nil {
			return nil, err
		}

		imageURLs = append(imageURLs, imageURL)
	}

	return imageURLs, nil
}

func (s *ProductService) GetAllProducts() ([]models.Products, error) {
	return s.ProductRepository.GetAllProducts()
}
