package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

type Transaction struct {
	Id               int `gorm:"primary_key"`
	Transaction_type string
	Amount           float64
	Category         string
	Description      string
	Date             string
}

func NewTransaction(transaction_type, category, description, date string, amount float64) (Transaction, error) {
	transaction := Transaction{
		Transaction_type: transaction_type,
		Amount:           amount,
		Category:         category,
		Description:      description,
		Date:             date,
	}
	if err := db.Create(transaction).Error; err != nil {
		return transaction, errors.New("transaction created fail")
	}
	return transaction, nil
}

func UserCreateTransaction(all_transaction []Transaction, idCounter *int, reader *bufio.Reader) ([]Transaction, error) {
	//id := *idCounter
	fmt.Println("Введите тип операции")
	transaction_type, _ := reader.ReadString('\n')
	transaction_type = strings.TrimSpace(transaction_type)
	if transaction_type == "Доход" || transaction_type == "Расход" {
		fmt.Println("Введите категорию")
		category, _ := reader.ReadString('\n')
		category = strings.TrimSpace(category)
		fmt.Println("Введите описание операции")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)
		fmt.Println("Введите сумму операции")
		var amount float64
		fmt.Scan(&amount)
		if amount <= 0 {
			return all_transaction, errors.New("Введена некорректная сумма операции\n")
		}
		transaction, _ := NewTransaction(transaction_type, category, description, time.Now().Format(time.DateTime), math.Abs(amount))
		all_transaction = append(all_transaction, transaction)
		*idCounter++
		return all_transaction, nil
	}
	return all_transaction, errors.New("Введен некорректный тип операции, попробуйте снова\n")
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
