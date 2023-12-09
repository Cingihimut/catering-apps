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
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()

	sellerRepository := repositories.NewSellerRepository(appConfig.DB)
	sellerService := services.NewSellerService(*sellerRepository)
	sellerController := controllers.NewSellerController(sellerService)
	routes.InitSellerRoutes(appConfig.App, sellerController)

	userRepository := repositories.NewUserRepository(appConfig.DB)
	userService := services.NewUserService(*userRepository)
	userController := controllers.NewUserController(userService)
	routes.InitUserRoutes(appConfig.App, userController)

	serverAddress := fmt.Sprintf(":%d", appConfig.Port)
	log.Printf("Server is running on %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, appConfig.App)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
