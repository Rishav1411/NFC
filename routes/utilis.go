package routes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneValidor(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	regex := regexp.MustCompile(`^\+\d{12}$`)
	return regex.Match([]byte(phone))
}

func WriteJson(w http.ResponseWriter, message []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(message)
}

func ServerError(w http.ResponseWriter) {
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "server error",
	})
	WriteJson(w, jsonData, http.StatusInternalServerError)
}
