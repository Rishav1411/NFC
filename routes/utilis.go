package routes

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func PhoneValidator(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	regex := regexp.MustCompile(`^\+\d{12}$`)
	return regex.Match([]byte(phone))
}
func OTPValidator(f1 validator.FieldLevel) bool {
	otp := f1.Field().String()
	regex := regexp.MustCompile(`^\d{4}$`)
	return regex.Match([]byte(otp))
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
	var TWILIO_PH_NO string = os.Getenv("TWILIO_PH_NO")

	var TWILIO_SID string = os.Getenv("TWILIO_SID")

	var TWILIO_TOKEN string = os.Getenv("TWILIO_TOKEN")
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

func GenerateToken(id int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().AddDate(0, 3, 0).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(KEY)
	jwt, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
func VerifyJWT(token string) int64 {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("THERE WAS ERROR IN PARSING")
		}
		return []byte(KEY), nil
	})
	if err != nil || !parsedToken.Valid {
		return -1
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1
	}
	if exp := claims["exp"].(float64); int(exp) < int(time.Now().Local().Unix()) {
		return -1
	}
	return int64(claims["user_id"].(float64))
}
