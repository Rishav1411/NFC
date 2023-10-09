package routes

import (
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
