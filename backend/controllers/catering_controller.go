package controllers

import (
	"net/http"
	"strconv"

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
	sellerId, exist := ctx.Get("id")
	if !exist{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Seller Id not found",
		})
        return
	}
	priceStr := ctx.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return 
	}
	catering.SellerID  = sellerId.(uint)
	catering.CateringName = ctx.PostForm("catering_name")
	catering.Description = ctx.PostForm("description")
	catering.Price = price
	form, err := ctx.MultipartForm()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return 
	}

	files := form.File["images"]
	imageURLs, err := c.CateringService.SaveImages(files)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	createdCatering, err := c.CateringService.Create(&catering, imageURLs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   createdCatering,
	})
}

func (c *CateringController) GetAllCatering(ctx *gin.Context) {
	caterings, err := c.CateringService.GetAllCatering()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get catering",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   caterings,
	})
}

func (c *CateringController) GetCateringBySellerID(ctx *gin.Context) {

	sellerId, exist := ctx.Get("id")
	if !exist{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Seller Id not found",
		})
        return
	}
	uintId := sellerId.(uint)
	
	catering, err := c.CateringService.GetCateringBySellerID(uint(uintId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to parse ID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   catering,
	})
}
