package main

import (
	"errors"
	"time"
)

type Transaction struct {
	Id               int     `json:"id" gorm:"primary_key"`
	Transaction_type string  `json:"type" binding:"required,oneof=income expense"`
	Amount           float64 `json:"amount" binding:"required,gt=0"`
	Category         string  `json:"category" binding:"required"`
	Description      string  `json:"description"`
	Date             int64   `json:"date" binding:"max=2147483647"`
}

func NewTransaction(transaction_type, category, description string, amount float64, date int64) error {
	transaction := Transaction{
		Transaction_type: transaction_type,
		Amount:           amount,
		Category:         category,
		Description:      description,
		Date:             date,
	}
	if date == 0 {
		transaction.Date = time.Now().Unix()
	}
	if err := db.Create(&transaction).Error; err != nil {
		return errors.New("transaction created fail")
	}
	return nil
}

func remove(id int, all_transactiom []Transaction) ([]Transaction, error) {
	if len(all_transactiom) == 0 || len(all_transactiom) < id {
		return all_transactiom, errors.New("Введен некорректный Id операции, попробуйте снова\n")
	}
	i := 0
	for idx, transaction := range all_transactiom {
		if transaction.Id != id {
			all_transactiom[i] = all_transactiom[idx]
			i++
		}
	}
	if id == len(all_transactiom) {
		return all_transactiom, errors.New("Введен некорректный Id операции, попробуйте снова\n")
	}
	return all_transactiom[:i], nil
}
