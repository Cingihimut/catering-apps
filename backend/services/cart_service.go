package services

import (
	"github.com/Cingihimut/catering-apps/models/converter"
	"github.com/Cingihimut/catering-apps/repositories"
)

type CartService struct {
	CartRepository repositories.CartRepository
}

func NewCartService(cartRepo repositories.CartRepository) *CartService {
	return &CartService{
		CartRepository: cartRepo,
	}
}

func (s *CartService) AddProduct(userID, productID, quantity uint) error {
	return s.CartRepository.AddToCart(userID, productID, quantity)
}

func (s *CartService) GetUserCart(userID uint) ([]converter.UserCartProduct, error) {
	return s.CartRepository.GetUserCart(userID)
}
