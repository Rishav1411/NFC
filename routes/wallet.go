package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"nfc/m/database"
	"nfc/m/database/operations"

	"github.com/go-chi/chi/v5"
)

type ContextKey string

const userContextKey ContextKey = "user_id"

func createWallet(w http.ResponseWriter, r *http.Request) {
	db := database.SQLConnection()
	if db == nil {
		ServerError(w)
		return
	}
	defer db.Close()
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		ServerError(w)
		return
	}
	defer tx.Rollback()
	id := operations.CreateWallet(r.Context().Value(userContextKey).(int64), tx)
	if id == -2 {
		ServerError(w)
		return
	}
	tx.Commit()
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": id,
	})
	WriteJson(w, jsonData, 201)
}

func Wallet() *chi.Mux {
	wallet := chi.NewRouter()
	wallet.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				jsonData, _ := json.Marshal(map[string]interface{}{
					"details": "not Authorization",
				})
				WriteJson(w, jsonData, http.StatusUnauthorized)
				return
			}
			user_id := VerifyJWT(token)
			if user_id == -1 {
				jsonData, _ := json.Marshal(map[string]interface{}{
					"details": "not Authorization",
				})
				WriteJson(w, jsonData, http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userContextKey, user_id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	wallet.Get("/", createWallet)
	wallet.Post("/transfer", transfer)
	return wallet
}
