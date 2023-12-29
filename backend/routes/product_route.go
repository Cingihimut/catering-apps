package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.Engine, productController *controllers.ProductController) {

	router.GET("/api/products", productController.GetAllProducts)
	router.POST("/api/products", middlewares.AuthMiddleware(), productController.Create)
	router.PUT("/api/products/:id", middlewares.AuthMiddleware(), productController.Update)
	router.DELETE("/api/products/:id", middlewares.AuthMiddleware(), productController.Delete)

}
