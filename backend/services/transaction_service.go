package services

import (
	"github.com/Cingihimut/catering-apps/models"
	"github.com/Cingihimut/catering-apps/repositories"
)

type TransactionService struct {
	TransactionRepo repositories.TransactionRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository) *TransactionService {
	return &TransactionService{
		TransactionRepo: transactionRepo,
	}
}

func (s *TransactionService) GetAllTransactions() ([]models.Transactions, error) {
	return s.TransactionRepo.GetAllTransactions()
}

func (s *TransactionService) CreateTransactionFromCart(buyerID, addressID uint, total float64) error {
	return s.TransactionRepo.CreateTransactionFromCart(buyerID, addressID, total)
}

func (s *TransactionService) GetTransactionsByUserID(userID uint) ([]models.Transactions, error) {
	return s.TransactionRepo.GetTransactionsByUserID(userID)
}
