package main

import (
	"encoding/json"
	"net/http"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		allTransaction, err := SelectAllTransaction()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allTransaction)
	case http.MethodPost:
		var req Transaction
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if err := db.Create(&req).Error; err != nil {
			http.Error(w, "colud not save transaction", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
