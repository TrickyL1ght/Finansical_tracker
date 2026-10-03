package main

import (
	"errors"
)

func SelectAllTransaction() ([]Transaction, error) {
	var allTransactions []Transaction
	if err := db.Find(&allTransactions).Error; err != nil {
		return nil, err
	}
	if len(allTransactions) == 0 {
		return nil, errors.New("Transaction not found")
	}
	return allTransactions, nil
}

func SelectTransactionById(id int) (Transaction, error) {
	var transaction Transaction
	if err := db.First(&transaction, id).Error; err != nil {
		return Transaction{}, errors.New("Transaction not found")
	}
	return transaction, nil
}
