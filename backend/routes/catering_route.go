package routes

import (
<<<<<<< HEAD
=======

>>>>>>> bb3f6c0e49cb4cfef00dda82bae5d65965057c5f
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCateringRoutes(router *gin.Engine, cateringController *controllers.CateringController) {
	

	router.GET("/api/caterings", cateringController.GetAll)
<<<<<<< HEAD
	router.GET("/api/caterings:sellerId", cateringController.GetCateringBySellerID)
=======
	router.GET("/api/caterings/:sellerId", cateringController.GetCateringBySellerID)
>>>>>>> bb3f6c0e49cb4cfef00dda82bae5d65965057c5f
    router.POST("/api/caterings",middlewares.AuthMiddleware(), cateringController.Create)

   
}