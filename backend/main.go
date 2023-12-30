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
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()

	// Allow CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	appConfig.App.Use(cors.New(config))

	// Serve Images
	appConfig.App.Static("/uploads/", "./uploads")

	// user route
	userRepository := repositories.NewUserRepository(appConfig.DB)
	userService := services.NewUserService(*userRepository)
	userController := controllers.NewUserController(userService)
	routes.InitUserRoutes(appConfig.App, userController)

	// Product
	productRepository := repositories.NewProductRepository(appConfig.DB)
	productService := services.NewProductService(*productRepository)
	productController := controllers.NewProductController(productService)
	routes.InitProductRoutes(appConfig.App, productController)

	// Categories
	categoriesRepository := repositories.NewCategoryRepository(appConfig.DB)
	categoriesService := services.NewCategoryService(*categoriesRepository)
	categoriesController := controllers.NewCategoryController(categoriesService)
	routes.InitCategoryRoutes(appConfig.App, categoriesController)

	// Cart
	cartRepository := repositories.NewCartRepository(appConfig.DB)
	cartService := services.NewCartService(*cartRepository)
	cartController := controllers.NewCartController(cartService)
	routes.InitCartRoutes(appConfig.App, cartController)

	// Transaction
	transactionRepository := repositories.NewTransactionRepository(appConfig.DB)
	transactionService := services.NewTransactionService(*transactionRepository)
	transactionController := controllers.NewTransactionController(transactionService)
	routes.InitTransactionRoutes(appConfig.App, transactionController)

	// running server
	serverAddress := fmt.Sprintf(":%d", appConfig.Port)
	log.Printf("Server is running on %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, appConfig.App)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
