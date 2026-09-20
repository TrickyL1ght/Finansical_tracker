package main

import (
	"errors"
	"fmt"
	"time"
)

func calculateBalance(all_transaction []Transaction) {
	var balance float64
	for _, transaction := range all_transaction {
		if transaction.Transaction_type == "Доход" {
			balance += transaction.Amount
		} else {
			balance -= transaction.Amount
		}
	}
	fmt.Printf("Ваш баланс составляет:%.2f\n\n", balance)
}

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

func SelectTransactionOfType(all_transaction []Transaction, transaction_type string) {
	for _, transaction := range all_transaction {
		if transaction.Transaction_type == transaction_type {
			fmt.Printf("Операция №%d\nСумма:%.2f\nТип операции:%s\nОписание:%s\nКатегория:%s\nДата:%s\n\n",
				transaction.Id, transaction.Amount, transaction.Transaction_type, transaction.Description, transaction.Category, transaction.Date)
		}
	}
}

func SelectTransactionOfCategory(all_transaction []Transaction, category string) {
	transactionCounter := 0
	for _, transaction := range all_transaction {
		if transaction.Category == category {
			fmt.Printf("Операция №%d\nСумма:%.2f\nТип операции:%s\nОписание:%s\nКатегория:%s\nДата:%s\n\n",
				transaction.Id, transaction.Amount, transaction.Transaction_type, transaction.Description, transaction.Category, transaction.Date)
			transactionCounter++
		}
	}
	if transactionCounter == 0 {
		fmt.Println("Операции по введенной категории не найдены")
	}
}

func StatisticByCategory(all_transaction []Transaction) {
	stat := make(map[string]float64)
	for _, transaction := range all_transaction {
		if stat[transaction.Category] != 0 {
			stat[transaction.Category] += transaction.Amount
		}
		stat[transaction.Category] = transaction.Amount
	}
	for key, value := range stat {
		fmt.Printf("Категория:%s - %.2f руб.", key, value)
	}
}

func TimeBasedOperations(userChoise string, all_transaction []Transaction) {
	for _, transaction := range all_transaction {
		a, _ := time.Parse(time.DateTime, transaction.Date)
		fmt.Println(a)
		fmt.Printf("%T\n", a)
	}
}
