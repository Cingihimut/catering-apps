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
	createdCatering, err := s.CateringRepository.Create(catering)
	if err != nil {
		return nil, err
	}
	var cateringImage models.CateringImages
	for _, imageURL := range imageURLs {
		
		cateringImage.CateringID = createdCatering.ID
		cateringImage.ImageURL = imageURL
		
		_, err := s.CateringRepository.CreateImage( &cateringImage)
		if err != nil {
			
			return nil, err
		}
	}

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
	caterings, err := s.CateringRepository.GetAllCatering()
	if err != nil {
		return nil, err
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
