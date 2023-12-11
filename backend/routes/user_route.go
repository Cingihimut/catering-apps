package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, userController *controllers.UserController) {

	r.GET("/", userController.Create)
}
