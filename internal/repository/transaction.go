package repository

import (
	"awesomeProject/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type TransactionGORM struct {
	db *gorm.DB
}

func NewTransactionGORM(db *gorm.DB) *TransactionGORM {
	return &TransactionGORM{db: db}
}

func (r *TransactionGORM) GetAllTransaction() ([]model.Transaction, error) {
	var allTransactions []model.Transaction
	if err := r.db.Find(&allTransactions).Error; err != nil {
		return nil, err
	}
	return allTransactions, nil
}

func (r *TransactionGORM) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return model.Transaction{}, fmt.Errorf("transaction created fail: %w", err)
	}
	return transaction, nil
}

func (r *TransactionGORM) UpdateTransaction(transaction model.Transaction) (model.Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return model.Transaction{}, fmt.Errorf("transaction updated fail: %w", err)
	}
	return transaction, nil
}

func (r *TransactionGORM) DeleteTransaction(id int) error {
	if err := r.db.Delete(&model.Transaction{}, id).Error; err != nil {
		return fmt.Errorf("transaction deleted fail: %w", err)
	}
	return nil
}

func (r *TransactionGORM) GetTransactionById(id int) (model.Transaction, error) {
	var transaction model.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return model.Transaction{}, fmt.Errorf("Transaction not found: %w", err)
	}
	return transaction, nil
}
