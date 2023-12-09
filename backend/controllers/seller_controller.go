// controllers/Seller_controller.go

package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/models/converter"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-gonic/gin"
)

type SellerController struct {
	SellerService *services.SellerService
}

func NewSellerController(SellerService *services.SellerService) *SellerController {
	return &SellerController{
		SellerService: SellerService,
	}
}

func (c *SellerController) Create(ctx *gin.Context) {
	var seller models.Sellers

	if err := ctx.ShouldBindJSON(&seller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	createdSeller, err := c.SellerService.CreateSeller(&seller)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	sellerResponse := converter.ConvertToSellerResponse(createdSeller)
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   sellerResponse,
	})
}

func (c *SellerController) Login(ctx *gin.Context) {
	var seller models.Sellers

	if err := ctx.ShouldBindJSON(&seller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	existingSeller, err := c.SellerService.LoginSeller(seller.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	if existingSeller.Password != seller.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid password",
		})
		return
	}

	token, err := c.SellerService.GenerateToken(existingSeller)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"token":  token,
	})
}
