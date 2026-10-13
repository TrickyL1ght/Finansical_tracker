package repository

import (
	"awesomeProject/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	TransactionRepository
	UserRepository
}

type TransactionRepository interface {
	GetAllTransaction(userId int) ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	UpdateTransaction(transaction model.Transaction) (model.Transaction, error)
	DeleteTransaction(id int) error
	GetTransactionById(id int) (model.Transaction, error)
}

type UserRepository interface {
	CrateUser(user model.User) error
	GetUser(login, password string) (model.User, error)
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		TransactionRepository: NewTransactionGORM(db),
		UserRepository:        NewUserGORM(db),
	}
}
