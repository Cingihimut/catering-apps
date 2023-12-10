package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-gonic/gin"
)

type CateringController struct {
	CateringService *services.CateringService
}

func NewCateringController(CateringService *services.CateringService) *CateringController {
	return &CateringController{
		CateringService: CateringService,
	}
}

func (c *CateringController) Create(ctx *gin.Context) {
	var catering models.Caterings
	id, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "sellerID not found in context"})
		return
	}

	floatID, ok := id.(float64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to convert sellerID to float64",
		})
		return
	}

	uintID := uint(floatID)


	catering.SellerID = uintID

	images := ctx.Request.MultipartForm.File["images"]
	if err := ctx.ShouldBindJSON(&catering); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var imageURLs []string
	if err := ctx.ShouldBindJSON(&imageURLs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	createdCatering, err := c.CateringService.Create(&catering, images)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   createdCatering,
	})
}
