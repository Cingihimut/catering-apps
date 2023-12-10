package services

import (
	"os"
	"time"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(UserRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: UserRepository,
	}
}

func (u *UserService) CreateUser(user *models.Users) (*models.Users, error) {

	createdUser, err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UserService) LoginUser(email string) (*models.Users, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GenerateToken(user *models.Users) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
