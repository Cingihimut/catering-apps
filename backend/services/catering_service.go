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

func (s *CateringService) Create( catering *models.Caterings, imageURLs []string) (*models.Caterings, error) {
	tx := s.CateringRepository.DB.Begin()
	createdCatering, err := s.CateringRepository.Create(tx,catering)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, imageURL := range imageURLs {
		cateringImage := models.CateringImages{
			CateringID: createdCatering.ID,
			ImageURL:   imageURL,
		}
	
		if _, err := s.CateringRepository.SaveImage(tx, &cateringImage); err != nil {
			tx.Rollback()
			return nil, err
		}
		createdCatering.Images = append(createdCatering.Images, cateringImage)

	}
	tx.Commit()

	return createdCatering, nil
}



func (s *CateringService) SaveImages(files []*multipart.FileHeader) ([]string, error) {
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


func (s *CateringService) GetAllCatering() ([]models.Caterings, error) {
	caterings, err := s.CateringRepository.GetAllCateringFromDB()
    if err != nil {
        return nil, err
    }

    for i := range caterings {
        s.CateringRepository.LoadImages(&caterings[i])
    }

    return caterings, nil
}



func (s *CateringService) GetCateringBySellerID(cateringID uint) (*models.Caterings, error) {
	catering, err := s.CateringRepository.GetCateringByID(cateringID)
	if err != nil {
		return nil, err
	}
	return catering, nil
}
