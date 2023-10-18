package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ContextKey string

const userContextKey ContextKey = "user_id"

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
	return wallet
}
