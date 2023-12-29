package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCartRoutes(router *gin.Engine, cartController *controllers.CartController) {

	router.GET("/api/cart", middlewares.AuthMiddleware(), cartController.GetUserCart)
	router.POST("/api/cart", middlewares.AuthMiddleware(), cartController.AddProduct)

}
