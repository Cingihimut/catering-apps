package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/gin-gonic/gin"
)

func InitSellerRoutes(router *gin.Engine, sellerController *controllers.SellerController) {
	sellerGroup := router.Group("/api/sellers")

	sellerGroup.POST("/register", sellerController.Create)
	sellerGroup.POST("/login", sellerController.Login)
}
