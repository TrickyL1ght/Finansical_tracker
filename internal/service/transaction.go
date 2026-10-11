package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
	"errors"
	"time"
)

type TransactionService struct {
	repository repository.Repository
}

func NewTransactionService(repository repository.Repository) *TransactionService {
	return &TransactionService{repository: repository}
}

func (s *TransactionService) GetAllTransaction() ([]model.Transaction, error) {
	transactions, err := s.repository.GetAllTransaction()
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, errors.New("no transactions found")
	}

	return transactions, nil
}

func (s *TransactionService) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {

	if transaction.Date == 0 {
		transaction.Date = time.Now().Unix()
	}

	return s.repository.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateTransaction(transaction model.Transaction) (model.Transaction, error) {

	if transaction.Date == 0 {
		transaction.Date = time.Now().Unix()
	}

	return s.repository.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(id int) error {
	return s.repository.DeleteTransaction(id)
}

func (s *TransactionService) GetTransactionById(id int) (model.Transaction, error) {
	return s.repository.GetTransactionById(id)
}
