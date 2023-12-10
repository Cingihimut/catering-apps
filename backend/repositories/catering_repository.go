package repositories

import (
	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/gorm"
)
type CateringRepository struct {
	DB *gorm.DB
}

func NewCateringRepository(db *gorm.DB) *CateringRepository {
	return &CateringRepository{
		DB: db,
	}
}


func (c *CateringRepository) Create(tx *gorm.DB, catering *models.Caterings) (*models.Caterings, error){
	query := "INSERT INTO CATERING (seller_id, catering_name, desciption, price)"
	result := tx.Exec(query, catering.SellerID, catering.CateringName, catering.Description, catering.Price)

	if result.Error != nil {
		return nil, result.Error
	}

	return catering, nil
}

func (c *CateringRepository) CreateImage(tx *gorm.DB, cateringImages *models.CateringImages) (*models.CateringImages, error){
	query := "INSERT INTO CATERING_IMAGES (catering_id, image_url)"
	result := tx.Exec(query, cateringImages.CateringID, cateringImages.ImageURL)

	if result.Error != nil {
		return nil, result.Error
	}

	return cateringImages, nil
}

func (c *CateringRepository) BeginTransaction()  *gorm.DB{
	query := "BEGIN"
	c.DB.Exec(query)

	return c.DB
}

