package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/Cingihimut/catering-apps/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(ProductService *services.ProductService) *ProductController {
	return &ProductController{
		ProductService: ProductService,
	}
}

func (c *ProductController) Create(ctx *gin.Context) {
	var product models.Products
	priceStr := ctx.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	product.ProductName = ctx.PostForm("product_name")
	product.Description = ctx.PostForm("description")
	product.Price = price
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	files := form.File["images"]
	imageURLs, err := c.ProductService.SaveImages(files)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}
	fmt.Println(form)
	fmt.Println(imageURLs)
	createdProduct, err := c.ProductService.Create(&product, imageURLs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   createdProduct,
	})
}

func (c *ProductController) GetAll(ctx *gin.Context) {
	products, err := c.ProductService.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get product",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

func (c *ProductController) GetProductBySellerID(ctx *gin.Context) {

	userId := ctx.Param("sellerId")
	sellerId, err := utils.ParseStrParamsToUint(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid user ID"})
		return
	}
	product, err := c.ProductService.GetProductBySellerID(sellerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to parse ID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}
