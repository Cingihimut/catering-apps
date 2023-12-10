package repositories

import (
	"fmt"

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


func (c *CateringRepository) Create( catering *models.Caterings) (*models.Caterings, error){
	query := "INSERT INTO caterings (seller_id, catering_name, description, price) VALUES (?, ?, ?, ?) RETURNING id"

	resultCatering:= c.DB.Exec(query, catering.SellerID, catering.CateringName, catering.Description, catering.Price)
	
	if resultCatering.Error != nil {
		return nil, resultCatering.Error
	}

	return catering, nil
}

func (c *CateringRepository) CreateImage(cateringImages *models.CateringImages) (*models.CateringImages, error){
	query := "INSERT INTO CATERING_IMAGES (catering_id, image_url) VALUES ( ? , ?)"
	result := c.DB.Exec(query, cateringImages.CateringID, cateringImages.ImageURL)

	if result.Error != nil {
		return nil, result.Error
	}
	return cateringImages, nil
}


func (c *CateringRepository) GetAllCatering() ([]models.Caterings, error) {
	var caterings []models.Caterings

	result := c.DB.Raw("SELECT * FROM caterings").Scan(&caterings)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return caterings, nil
}

func (r *CateringRepository) GetCateringByID(cateringID uint) (*models.Caterings, error) {
	var catering models.Caterings

	result := r.DB.Raw("SELECT * FROM caterings WHERE seller_id = ?", cateringID).Scan(&catering)

	if result.Error != nil {
		return nil, result.Error
	}
	return &catering, nil
}