package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCategoryRoutes(router *gin.Engine, categoryController *controllers.CategoryController) {
	categoryRoutes := router.Group("/api/categories")
	{
		categoryRoutes.POST("/", middlewares.AuthMiddleware(), categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", middlewares.AuthMiddleware(), categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares.AuthMiddleware(), categoryController.DeleteCategory)
	}
}
