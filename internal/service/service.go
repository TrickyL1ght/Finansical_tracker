package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
)

type Service struct {
	User
	Transaction
}

type User interface {
	CreateUser(user model.User) error
	UserAuth(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Transaction interface {
	GetAllTransaction() ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	UpdateTransaction(transaction model.Transaction) (model.Transaction, error)
	DeleteTransaction(id int) error
	GetTransactionById(id int) (model.Transaction, error)
}

func NewService(repository repository.Repository) Service {
	return Service{
		User:        NewUserService(repository),
		Transaction: NewTransactionService(repository),
	}
}
