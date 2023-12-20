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

func (p *ProductService) Update(product *models.Products, imageURLs []string) (*models.Products, error) {
	tx := p.ProductRepository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	updatedProduct, err := p.ProductRepository.Update(tx, product)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := p.ProductRepository.SaveCategories(tx, updatedProduct.ID, product.Categories); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Hapus semua gambar produk lama sebelum menyimpan gambar yang baru
	if err := p.ProductRepository.DeleteImages(tx, updatedProduct.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, imageURL := range imageURLs {
		productImage := models.ProductImages{
			ProductID: updatedProduct.ID,
			ImageURL:  imageURL,
		}

		if _, err := p.ProductRepository.SaveImage(tx, &productImage); err != nil {
			tx.Rollback()
			return nil, err
		}
		updatedProduct.Images = append(updatedProduct.Images, productImage)
	}

	tx.Commit()

	return updatedProduct, nil
}



func (s *ProductService) Delete(productID uint) error {
	if err := s.ProductRepository.DeleteImagesByProductID(productID); err != nil {
        return err
    }

    // Menghapus produk
    return s.ProductRepository.Delete(productID)
}


func (s *ProductService) GetAllProducts() ([]models.Products, error) {
	return s.ProductRepository.GetAllProductsFromViews()
}
