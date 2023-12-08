package main

import (
	"os"

	"github.com/Cingihimut/catering-apps/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.InitDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
