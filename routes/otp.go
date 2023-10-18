package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"nfc/m/database"
	"nfc/m/database/operations"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func otp_validation(otpData []byte) *OTP {
	otp := &OTP{}
	otperr := json.Unmarshal(otpData, otp)
	if otperr != nil {
		return nil
	}
	validate := validator.New()
	validate.RegisterValidation("phone", PhoneValidator)
	validate.RegisterValidation("otp", OTPValidator)
	err := validate.Struct(validate)
	if err != nil {
		return nil
	}
	return otp
}

func Otp() *chi.Mux {
	otp := chi.NewRouter()
	otp.Post("/", func(w http.ResponseWriter, r *http.Request) {
		requestData, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		otp := otp_validation(requestData)
		if otp == nil {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "data is not of valid format",
			})
			WriteJson(w, jsonData, 400)
			return
		}
		client := database.RedisConnection()
		if client == nil {
			ServerError(w)
			return
		}
		defer client.Close()
		key := GenerateKey(otp.Phone)
		storedOtp, err := client.HGetAll(context.Background(), key).Result()
		if err != nil {
			ServerError(w)
			return
		}
		if len(storedOtp) == 0 {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "otp is expired",
			})
			WriteJson(w, jsonData, 403)
			return
		}

		if storedOtp["otp"] != otp.Otp {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "otp is not valid",
			})
			WriteJson(w, jsonData, 403)
			return
		}
		db := database.SQLConnection()
		if db == nil {
			ServerError(w)
			return
		}
		defer db.Close()
		ctx := context.Background()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			ServerError(w)
			return
		}
		defer tx.Rollback()
		var id int64
		if storedOtp["type"] == "sign_up" {
			id = operations.RegisterUser(storedOtp["phone"], storedOtp["name"], storedOtp["reg"], tx)
		} else if storedOtp["type"] == "login" {
			id = operations.CheckUser(otp.Phone, db)
		}
		if id == -2 {
			ServerError(w)
			return
		}
		jwt, err := GenerateToken(id)
		if err != nil {
			ServerError(w)
			return
		}
		err = tx.Commit()
		if err != nil {
			ServerError(w)
			return
		}
		jsonData, _ := json.Marshal(map[string]interface{}{
			"token": jwt,
			"type":  "Bearer",
		})
		if storedOtp["type"] == "sign_up" {
			WriteJson(w, jsonData, 201)
		} else {
			WriteJson(w, jsonData, 200)
		}

	})
	return otp
}
