package services

import (
	"os"
	"time"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type SellerService struct {
	SellerRepository repositories.SellerRepository
}

func NewSellerService(SellerRepository repositories.SellerRepository) *SellerService {
	return &SellerService{
		SellerRepository: SellerRepository,
	}
}

func (s *SellerService) CreateSeller(seller *models.Sellers) (*models.Sellers, error) {

	createdSeller, err := s.SellerRepository.Create(seller)
	if err != nil {
		return nil, err
	}

	return createdSeller, nil
}

func (s *SellerService) LoginSeller(email string) (*models.Sellers, error) {
	seller, err := s.SellerRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *SellerService) GenerateToken(seller *models.Sellers) (string, error) {
	claims := jwt.MapClaims{
		"id":    seller.ID,
		"email": seller.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
