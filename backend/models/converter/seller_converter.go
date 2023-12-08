package converter

import (
	"time"

	"github.com/Cingihimut/catering-apps/models"
)
type SellerResponse struct {
	ID             uint   `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ConvertToSellerResponse(seller *models.Sellers) *SellerResponse {
	return &SellerResponse{
		ID:            seller.ID,
		Email:         seller.Email,
		Username:      seller.UserName,
		CreatedAt:     seller.CreatedAt,
		UpdatedAt:     seller.UpdatedAt,
	}
}
