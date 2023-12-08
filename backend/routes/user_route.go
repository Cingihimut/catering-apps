package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userController := &controllers.User{}

	r.GET("/", userController.Get)
}
