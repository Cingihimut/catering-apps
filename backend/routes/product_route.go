package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.Engine, productController *controllers.ProductController) {

	router.GET("/api/products", productController.GetAllProducts)
	router.POST("/api/products", middlewares.AuthMiddleware(), productController.Create)

}
