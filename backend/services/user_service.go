package services

import (
	"errors"
	"os"
	"time"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := u.hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Role = models.RoleUser
	user.Password = hashedPassword
	createdUser, err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UserService) LoginUser(user *models.Users) (string, error) {
	existingUser, err := u.UserRepository.FindByEmail(user)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("Incorrect Password!")
	}

	token, err := u.GenerateToken(existingUser)
	if err != nil {
		return "", errors.New("Failed to generate Token")
	}

	return token, nil
}

func (u *UserService) GenerateToken(user *models.Users) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (u *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
