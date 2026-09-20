package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=127.0.0.1 user=postgres password=HomePass dbname=ftdb port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not init DB: %v", err)
	}

	if err = db.AutoMigrate(&Transaction{}); err != nil {
		log.Fatalf("could not migrate DB: %v", err)
	}
}

func main() {
	initDB()

	a := Transaction{
		1, "Доход", 10000, "ЗП", "Аванс", "2026-08-13 12:45:20",
	}
	if err := db.Create(&a).Error; err != nil {
		fmt.Errorf("Error: %s", err)
	}
	all_transaction := []Transaction{}
	idCounter := 2 //Уменьшить каунтер до 1, если не будет тестовых операций
	reader := bufio.NewReader(os.Stdin)
	var err error

	//Тестовые операции
	//all_transaction = append(all_transaction, Transaction{
	//	1, "Доход", 10000, "ЗП", "Аванс", time.Now().Format(time.DateTime),
	//})
	//all_transaction = append(all_transaction, Transaction{
	//	1, "Доход", 10000, "ЗП", "Аванс", "2026-08-13 12:45:20",
	//})

	for {
		fmt.Println("1. Добавить операцию\n2. Список операций\n3. Баланс\n4. Удаление операции\n5. Статистика по категориям\n6. Выход")
		userChoise, _ := reader.ReadString('\n')
		userChoise = strings.TrimSpace(userChoise)
		switch userChoise {
		case "1":
			all_transaction, err = UserCreateTransaction(all_transaction, &idCounter, reader)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Операция успешно создана\n\n")
			}
		case "2":
			fmt.Printf("Какие операции нужно вывести?\n1. Все операции\n2. Доходы\n3. Расходы\n4. Операции по категории\n")
			userChoise, err = reader.ReadString('\n')
			userChoise = strings.TrimSpace(userChoise)
			switch userChoise {
			case "1":
				selectAllTransaction(all_transaction)
			case "2":
				SelectTransactionOfType(all_transaction, "Доход")
			case "3":
				SelectTransactionOfType(all_transaction, "Расход")
			case "4":
				fmt.Println("Введите категорию")
				category, _ := reader.ReadString('\n')
				category = strings.TrimSpace(category)
				SelectTransactionOfCategory(all_transaction, category)
			case "5":
				fmt.Println("Введите дату (Формат ГГГГ-ММ-ДД)")
				userChoise, _ = reader.ReadString('\n')
				userChoise = strings.TrimSpace(userChoise)
				TimeBasedOperations(userChoise, all_transaction)
			default:
				fmt.Println("Введено некорректное значение")
			}

		case "3":
			calculateBalance(all_transaction)
		case "4":
			fmt.Println("Выберите операцию для удаления")
			var id int
			fmt.Scan(&id)
			all_transaction, err = remove(id, all_transaction)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Операция успешно удалена\n\n")
			}
		case "5":
			StatisticByCategory(all_transaction)
		case "6":
			fmt.Println("Выход из приложения")
			return
		default:
			fmt.Println("Введено некорректное значение")
		}
	}
}
