package services

import (
	"mime/multipart"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/Cingihimut/catering-apps/utils"
)

type CateringService struct {
	CateringRepository repositories.CateringRepository
}

func NewCateringService(CateringRepository repositories.CateringRepository) *CateringService {
	return &CateringService{
		CateringRepository: CateringRepository,
	}
}

func (s *CateringService) Create( catering *models.Caterings, files []*multipart.FileHeader) (*models.Caterings, error) {

	tx := s.CateringRepository.BeginTransaction()

	defer func() {
		if r := recover(); r != nil {
			tx.Exec("ROLLBACK")
		}
	}()

	catering,err := s.CateringRepository.Create(tx, catering)
	if err != nil {
		tx.Exec("ROLLBACK")
		return nil, err
	}

	imageURLs, err := s.saveCateringImages(catering.ID, files)
	if err != nil {
		tx.Exec("ROLLBACK")
		return nil, err
	}

	for _, imageURL := range imageURLs {
		cateringImages := models.CateringImages{
			CateringID: catering.ID,
			ImageURL:  imageURL,
		}

		if _,err := s.CateringRepository.CreateImage(tx, &cateringImages); err != nil {
			tx.Exec("ROLLBACK")
			return nil, err
		}
	}
	tx.Exec("COMMIT")
	return catering, nil
}





func (s *CateringService) saveCateringImages( cateringID uint, files []*multipart.FileHeader) ([]string, error) {
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

