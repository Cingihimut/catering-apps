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

func (c *UserController) Register(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	createdUser, err := c.UserService.CreateUser(&user)
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

func (c *UserController) RegisterOwner(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	createdUser, err := c.UserService.CreateOwner(&user)
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
func (c *UserController) Login(ctx *gin.Context) {
	var user models.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	token, err := c.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
