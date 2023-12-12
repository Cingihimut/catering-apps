package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Cingihimut/catering-apps/config"
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/repositories"
	"github.com/Cingihimut/catering-apps/routes"
	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()

	// Allow CORS
	appConfig.App.Use(cors.Default())
	// Handle Trailing Slah from grouping
	appConfig.App.Use(func(c *gin.Context) {
        if c.Request.URL.Path != "/" && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
            c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
            return
        }
        c.Next()
    })
	
	// seller route
	sellerRepository := repositories.NewSellerRepository(appConfig.DB)
	sellerService := services.NewSellerService(*sellerRepository)
	sellerController := controllers.NewSellerController(sellerService)
	routes.InitSellerRoutes(appConfig.App, sellerController)

	// user route
	userRepository := repositories.NewUserRepository(appConfig.DB)
	userService := services.NewUserService(*userRepository)
	userController := controllers.NewUserController(userService)
	routes.InitUserRoutes(appConfig.App, userController)

	// catering route
	cateringRepository := repositories.NewCateringRepository(appConfig.DB)
	cateringService := services.NewCateringService(*cateringRepository)
	cateringController := controllers.NewCateringController(cateringService)
	routes.InitCateringRoutes(appConfig.App, cateringController)

	// running server
	serverAddress := fmt.Sprintf(":%d", appConfig.Port)
	log.Printf("Server is running on %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, appConfig.App)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
