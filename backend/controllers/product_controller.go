package controllers

import (
	"log"
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
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Admin role required",
		})
		return
	}

	var product models.Products

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	log.Printf("Form Data: %v", ctx.Request.PostForm)

	price, err := strconv.ParseFloat(ctx.PostForm("price"), 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	product.Price = price
	categories := ctx.PostFormArray("categories")

	product.Categories = make([]models.Categories, len(categories))

	for i, categoryIDStr := range categories {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid category ID",
			})
			return
		}

		category := models.Categories{ID: uint(categoryID)}
		product.Categories[i] = category
	}

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
		return
	}

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

func (c *ProductController) Update(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Admin role required",
		})
		return
	}

	productIDStr := ctx.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid product ID",
		})
		return
	}

	var updatedProduct *models.Products
	if err := ctx.ShouldBind(&updatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	updatedProduct.ID = uint(productID)

	categories := ctx.PostFormArray("categories")

	updatedProduct.Categories = make([]models.Categories, len(categories))

	for i, categoryIDStr := range categories {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid category ID",
			})
			return
		}

		category := models.Categories{ID: uint(categoryID)}
		updatedProduct.Categories[i] = category
	}

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
		return
	}

	updatedProduct, err = c.ProductService.Update(updatedProduct, imageURLs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   updatedProduct,
	})
}

func (c *ProductController) Delete(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Admin role required",
		})
		return
	}
	productIDStr := ctx.Param("id")
	productID, err := utils.ParseStrParamsToUint(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid product ID",
		})
		return
	}
	err = c.ProductService.Delete(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.ProductService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}
