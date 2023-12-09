package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/api/users")

	userGroup.POST("/register", userController.Create)
	userGroup.POST("/login", userController.Login)
}
