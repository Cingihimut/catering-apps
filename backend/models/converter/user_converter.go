package converter

import (
	"time"

	"github.com/Cingihimut/catering-apps/models"
)

type UserResponse struct {
	Email     string    `json:"email"`
	Name  string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertToUserResponse(user *models.Users) *UserResponse {
	return &UserResponse{
		Email:     user.Email,
		Name:  user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
