package main

import (
	"os"

	"github.com/Cingihimut/catering-apps/config"
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.InitDB()
}

func main() {
	r := gin.Default()

	userController := &controllers.User{}

	r.GET("/", userController.Get)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
