package routes

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneValidor(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	regex := regexp.MustCompile(`^\+\d{12}$`)
	return regex.Match([]byte(phone))
}

func RegValidator(f1 validator.FieldLevel) bool {
	reg := f1.Field().String()
	regex := regexp.MustCompile(`^\d{2}[A-Z]{3}\d{4}$`)
	return regex.Match([]byte(reg))
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

func GenerateOTP() string {
	otp := fmt.Sprintf("%04d", rand.Intn(10000))
	return otp
}
func GenerateKey(phone string) string {
	hash := fnv.New32a()
	hash.Write([]byte(phone))
	hashValue := fmt.Sprintf("%d", hash.Sum32())
	return hashValue
}
