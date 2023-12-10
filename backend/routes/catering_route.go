package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCateringRoutes(router *gin.Engine, cateringController *controllers.CateringController) {
	cateringGroup := router.Group("/api/catering")

	// cateringGroup.GET("/", userController.CreateUser)
	{
		cateringGroup.Use(middlewares.AuthMiddleware())
		{
			cateringGroup.POST("/", cateringController.Create)
		}
	}
}
