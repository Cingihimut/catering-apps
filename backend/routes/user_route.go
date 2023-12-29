package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine, uc *controllers.UserController) {

	router.POST("/api/users/register", uc.Register)
	router.POST("/api/users/register-owner", middlewares.OwnerMiddleware(), uc.RegisterOwner)
	router.POST("/api/users/login", uc.Login)
}
