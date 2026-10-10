package transaction

import (
	"errors"
	"time"
)

type TransactionService interface {
	GetAllTransaction() ([]Transaction, error)
	CreateTransaction(transaction Transaction) (Transaction, error)
	UpdateTransaction(transaction Transaction) (Transaction, error)
	DeleteTransaction(id int) error
	GetTransactionById(id int) (Transaction, error)
}

type transactionService struct {
	repository TransactionRepository
}

func NewTransactionService(repository TransactionRepository) TransactionService {
	return &transactionService{repository: repository}
}

func (s *transactionService) GetAllTransaction() ([]Transaction, error) {
	transactions, err := s.repository.GetAllTransaction()
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, errors.New("no transactions found")
	}

	return transactions, nil
}

func (s *transactionService) CreateTransaction(transaction Transaction) (Transaction, error) {
	return s.repository.CreateTransaction(transaction)
}

func (s *transactionService) UpdateTransaction(transaction Transaction) (Transaction, error) {

	if transaction.Date == 0 {
		transaction.Date = time.Now().Unix()
	}

	return s.repository.UpdateTransaction(transaction)
}

func (s *transactionService) DeleteTransaction(id int) error {
	return s.repository.DeleteTransaction(id)
}

func (s *transactionService) GetTransactionById(id int) (Transaction, error) {
	return s.repository.GetTransactionById(id)
}
