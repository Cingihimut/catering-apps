package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCateringRoutes(router *gin.Engine, cateringController *controllers.CateringController) {
	

	router.GET("/api/caterings", cateringController.GetAll)
	router.GET("/api/caterings:sellerId", cateringController.GetCateringBySellerID)
    router.POST("/api/caterings",middlewares.AuthMiddleware(), cateringController.Create)

   
}