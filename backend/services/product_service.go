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

func (s *ProductService) Create(product *models.Products, imageURLs []string) (*models.Products, error) {
	tx := s.ProductRepository.DB.Begin()
	createdProduct, err := s.ProductRepository.Create(tx, product)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, imageURL := range imageURLs {
		productImage := models.ProductImages{
			ProductID: createdProduct.ID,
			ImageURL:  imageURL,
		}

		if _, err := s.ProductRepository.SaveImage(tx, &productImage); err != nil {
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

func (s *ProductService) GetAllProduct() ([]models.Products, error) {
	products, err := s.ProductRepository.GetAllProductFromDB()
	if err != nil {
		return nil, err
	}

	for i := range products {
		s.ProductRepository.LoadImages(&products[i])
	}

	return products, nil
}

func (s *ProductService) GetProductBySellerID(productID uint) (*models.Products, error) {
	product, err := s.ProductRepository.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
