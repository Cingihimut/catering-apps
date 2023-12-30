// transaction_repository.go

package repositories

import (
	"fmt"

	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) GetAllTransactions() ([]models.Transactions, error) {
	var transactions []models.Transactions
	query := "SELECT * FROM transactions"
	if err := r.DB.Raw(query).Scan(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) CreateTransactionFromCart(buyerID, addressID uint, total float64) error {
	// Implementasi pembuatan transaksi dari cart ke database di sini menggunakan RAW SQL
	// ...
	return nil
}

func (r *TransactionRepository) GetTransactionsByUserID(userID uint) ([]models.Transactions, error) {
	var transactions []models.Transactions
	query := fmt.Sprintf("SELECT * FROM transactions WHERE buyer_id = %d", userID)
	if err := r.DB.Raw(query).Scan(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
