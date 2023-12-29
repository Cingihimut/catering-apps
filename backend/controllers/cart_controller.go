package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/services"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService *services.CartService
}

func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{
		CartService: cartService,
	}
}

func (c *CartController) AddProduct(ctx *gin.Context) {
	var request struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  uint `json:"quantity" binding:"required"`
	}

	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	userID, ok := userId.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := c.CartService.AddProduct(userID, request.ProductID, request.Quantity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product added to cart successfully",
	})
}

func (c *CartController) GetUserCart(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	userID, ok := userId.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}

	userCart, err := c.CartService.GetUserCart(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   userCart,
	})
}
