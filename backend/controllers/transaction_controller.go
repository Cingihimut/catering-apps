package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/services"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService *services.TransactionService
}

func NewTransactionController(TransactionService *services.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: TransactionService,
	}
}

func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (c *TransactionController) CreateTransactionFromCart(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func (c *TransactionController) GetTransactionsByUserID(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	userID, ok := userId.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}
	transactions, err := c.TransactionService.GetTransactionsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}
