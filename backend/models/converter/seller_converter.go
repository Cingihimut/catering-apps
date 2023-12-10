package converter

import (
	"time"

	"github.com/Cingihimut/catering-apps/models"
)

type SellerResponse struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertToSellerResponse(seller *models.Sellers) *SellerResponse {
	return &SellerResponse{
		Email:     seller.Email,
		Username:  seller.UserName,
		CreatedAt: seller.CreatedAt,
		UpdatedAt: seller.UpdatedAt,
	}
}
