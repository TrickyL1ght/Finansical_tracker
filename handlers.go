package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		allTransaction, err := SelectAllTransaction()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			json.NewEncoder(w).Encode(err)
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		json.NewEncoder(w).Encode(allTransaction)

	case http.MethodPost:
		var req Transaction
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Transaction_type != "income" && req.Transaction_type != "expense" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			http.Error(w, "Invalid transaction type", http.StatusBadRequest)
			return
		}
		err := NewTransaction(req.Transaction_type, req.Category, req.Description, req.Amount)
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.WriteHeader(http.StatusNoContent)
	}
}

func EditTransactionHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/transactions/")
	transactionId := path
	switch r.Method {
	case http.MethodDelete:
		if err := db.Delete(&Transaction{}, "id = ?", transactionId).Error; err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			json.NewEncoder(w).Encode(err)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
