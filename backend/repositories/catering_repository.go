package repositories

import (
	"fmt"
	"log"

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
	query := "INSERT INTO caterings (seller_id, catering_name, description, price) VALUES (?, ?, ?, ?) RETURNING id"

	result:= tx.Raw(query,catering.SellerID, catering.CateringName, catering.Description, catering.Price).Row()
	
	if result.Err() != nil {
		return nil, result.Err()
	}

	var insertedID uint
	if err := result.Scan(&insertedID); err != nil {
		return nil, err
	}
	catering.ID = insertedID
	return catering, nil
}

func (c *CateringRepository) SaveImage(tx *gorm.DB, cateringImages *models.CateringImages) (*models.CateringImages, error){
    query := "INSERT INTO catering_images (catering_id, image_url) VALUES (?,?) RETURNING id"
    result := tx.Raw(query, cateringImages.CateringID, cateringImages.ImageURL).Scan(&cateringImages.ID)

    if result.Error != nil {
        log.Printf("Error executing query: %v\n", result.Error)
        return nil, result.Error
    }
    return cateringImages, nil
}


func (c *CateringRepository) GetAllCateringFromDB() ([]models.Caterings, error) {
	var caterings []models.Caterings
	result := c.DB.Raw("SELECT * FROM caterings").Scan(&caterings)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return caterings, nil
}

func (c *CateringRepository) LoadImages(catering *models.Caterings) {
    query := "SELECT * FROM catering_images WHERE catering_id = ?"
    result := c.DB.Raw(query, catering.ID).Scan(&catering.Images)

    if result.Error != nil {
        log.Printf("Error loading images: %v\n", result.Error)
    }
}


func (r *CateringRepository) GetCateringByID(cateringID uint) (*models.Caterings, error) {
	var catering models.Caterings

	result := r.DB.Raw("SELECT * FROM caterings WHERE seller_id = ?", cateringID).Scan(&catering)

	if result.Error != nil {
		return nil, result.Error
	}
	return &catering, nil
}