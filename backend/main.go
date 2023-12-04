package main

import (
	"github.com/Cingihimut/catering-apps/internal/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVar()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
