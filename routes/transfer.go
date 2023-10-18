package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"nfc/m/database"
	"nfc/m/database/operations"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func transaction_validation(transactionData []byte) *Transaction {
	trasaction := &Transaction{}
	err := json.Unmarshal(transactionData, trasaction)
	if err != nil {
		return nil
	}
	validator := validator.New()
	err = validator.Struct(trasaction)
	if err != nil {
		return nil
	}
	return trasaction
}

func transfer(w http.ResponseWriter, r *http.Request) {
	transactionData, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	transaction := transaction_validation(transactionData)
	if transaction == nil {
		ServerError(w)
		return
	}
	if r.Context().Value(userContextKey) != transaction.Sender_id && r.Context().Value(userContextKey) != transaction.Receiver_id {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "not Authorization",
		})
		WriteJson(w, jsonData, http.StatusUnauthorized)
		return
	}
	db := database.SQLConnection()
	if db == nil {
		ServerError(w)
		return
	}
	defer db.Close()
	sender_id, err := strconv.ParseInt(transaction.Sender_id, 10, 64)
	if err != nil {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "invalid id",
		})
		WriteJson(w, jsonData, http.StatusBadRequest)
		return
	}
	receiver_id, err := strconv.ParseInt(transaction.Receiver_id, 10, 64)
	if err != nil {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "invalid id",
		})
		WriteJson(w, jsonData, http.StatusBadRequest)
		return
	}
	balance := operations.CheckBalance(sender_id, db)
	if balance < int64(transaction.Amount) {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "balance is insufficient",
		})
		WriteJson(w, jsonData, http.StatusPaymentRequired)
		return
	}
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		ServerError(w)
		return
	}
	defer tx.Rollback()
	err = operations.UpdateBalance(sender_id, -int64(transaction.Amount), tx)
	if err != nil {
		ServerError(w)
		return
	}
	err = operations.UpdateBalance(receiver_id, int64(transaction.Amount), tx)
	if err != nil {
		ServerError(w)
		return
	}
	tx.Commit()
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "completed successfully",
	})
	WriteJson(w, jsonData, 200)
}
