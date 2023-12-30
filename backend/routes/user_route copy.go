package routes

import (
	"github.com/Cingihimut/catering-apps/controllers"
	"github.com/Cingihimut/catering-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func InitTransactionRoutes(router *gin.Engine, tc *controllers.TransactionController) {

	router.GET("/api/transaction", middlewares.AuthMiddleware(), tc.GetAllTransactions)
	router.POST("/api/transaction", tc.CreateTransactionFromCart)
	router.GET("/api/transaction/:userId", tc.GetTransactionsByUserID)
}
