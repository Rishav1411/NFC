package routes

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
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

func SendSMS(phone string, otp string) bool {
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: TWILIO_SID,
			Password: TWILIO_TOKEN,
		},
	)
	params := &api.CreateMessageParams{}
	params.SetBody(fmt.Sprintf("this is the otp code for verification %s which is valid only for 3 minutes", otp))
	params.SetFrom(TWILIO_PH_NO)
	params.SetTo(phone)
	_, err := client.Api.CreateMessage(params)
	return err == nil
}
