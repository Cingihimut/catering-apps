package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCategoryRoutes(router *gin.Engine, categoryController *controllers.CategoryController) {
	// router.POST("/api/categories", categoryController.GetAll())
	router.POST("/api/categories", middlewares.AuthMiddleware(), categoryController.CreateCategory)
	router.PUT("/api/categories/:id", middlewares.AuthMiddleware(), categoryController.UpdateCategory)
	router.DELETE("/api/categories/:id", middlewares.AuthMiddleware(), categoryController.DeleteCategory)
}
