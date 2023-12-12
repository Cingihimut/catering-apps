package routes

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCateringRoutes(router *gin.Engine, cateringController *controllers.CateringController) {
	
	cateringGroup := router.Group("/api/caterings")
	cateringGroup.Use(func(c *gin.Context) {
        if c.Request.URL.Path != "/" && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
            c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
            return
        }
        c.Next()
    })
	cateringGroup.GET("/", cateringController.GetAll)
	cateringGroup.GET("/:sellerId", cateringController.GetCateringBySellerID)
    cateringGroup.POST("/",middlewares.AuthMiddleware(), cateringController.Create)

   
}