package transaction

import (
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction() ([]Transaction, error)
	CreateTransaction(transaction Transaction) (Transaction, error)
	UpdateTransaction(transaction Transaction) (Transaction, error)
	DeleteTransaction(id int) error
	GetTransactionById(id int) (Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetAllTransaction() ([]Transaction, error) {
	var allTransactions []Transaction
	if err := r.db.Find(&allTransactions).Error; err != nil {
		return nil, err
	}
	return allTransactions, nil
}

func (r *transactionRepository) CreateTransaction(transaction Transaction) (Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return Transaction{}, fmt.Errorf("transaction created fail: %w", err)
	}
	return transaction, nil
}

func (r *transactionRepository) UpdateTransaction(transaction Transaction) (Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return Transaction{}, fmt.Errorf("transaction updated fail: %w", err)
	}
	return transaction, nil
}

func (r *transactionRepository) DeleteTransaction(id int) error {
	if err := r.db.Delete(&Transaction{}, id).Error; err != nil {
		return fmt.Errorf("transaction deleted fail: %w", err)
	}
	return nil
}

func (r *transactionRepository) GetTransactionById(id int) (Transaction, error) {
	var transaction Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return Transaction{}, fmt.Errorf("Transaction not found: %w", err)
	}
	return transaction, nil
}
