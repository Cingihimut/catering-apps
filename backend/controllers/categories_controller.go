package controllers

import (
	"net/http"
	"strconv"

	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Owner role required",
		})
		return
	}

	var category models.Categories
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_, err := c.CategoryService.Create(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   "Category Added !",
	})
}
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Owner role required",
		})
		return
	}

	categoryID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid category ID",
		})
		return
	}

	var updatedCategoryInput models.Categories
	if err := ctx.ShouldBindJSON(&updatedCategoryInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON format",
		})
		return
	}

	updatedCategoryInput.ID = uint(categoryID)
	updatedCategoryOutput, err := c.CategoryService.Update(&updatedCategoryInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   updatedCategoryOutput,
	})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || models.EnumRole(role.(string)) != models.RoleOwner {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized: Owner role required",
		})
		return
	}

	categoryID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid category ID",
		})
		return
	}

	if err := c.CategoryService.Delete(uint(categoryID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Category deleted successfully",
	})
}

func (c *CategoryController) GetAll(ctx *gin.Context) {

	result, err := c.CategoryService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}
