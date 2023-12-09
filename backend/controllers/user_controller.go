// controllers/Seller_controller.go

package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/models/converter"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var user models.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	createdUser, err := u.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	userResponse := converter.ConvertToUserResponse(createdUser)
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   userResponse,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var user models.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	existingUser, err := u.UserService.LoginUser(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	if existingUser.Password != user.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid password",
		})
		return
	}

	token, err := u.UserService.GenerateToken(existingUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"token":  token,
	})
}
