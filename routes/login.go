package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"nfc/m/database"
	"nfc/m/database/operations"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func phone_validation(jsonData []byte) *Phone {
	phone := &Phone{}
	phoneerr := json.Unmarshal(jsonData, phone)
	if phoneerr != nil {
		return nil
	}
	validate := validator.New()
	validate.RegisterValidation("phone", PhoneValidator)
	err := validate.Struct(phone)
	if err != nil {
		return nil
	}
	return phone
}

func Login() *chi.Mux {
	login := chi.NewMux()
	login.Post("/", func(w http.ResponseWriter, r *http.Request) {
		requestData, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		phone := phone_validation(requestData)
		if phone == nil {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "data is not of valid format",
			})
			WriteJson(w, jsonData, 400)
			return
		}
		db := database.SQLConnection()
		if db == nil {
			ServerError(w)
			return
		}
		defer db.Close()
		id := operations.CheckUser(phone.PhoneNumber, db)
		if id == -2 {
			ServerError(w)
			return
		}
		if id == -1 {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "user didnt exist",
			})
			WriteJson(w, jsonData, 404)
			return
		}
		client := database.RedisConnection()
		if client == nil {
			ServerError(w)
			return
		}
		defer client.Close()
		otp := GenerateOTP()
		key := GenerateKey(phone.PhoneNumber)
		if !SendSMS(phone.PhoneNumber, otp) {
			ServerError(w)
			return
		}
		data := map[string]interface{}{
			"type": "login",
			"otp":  otp,
		}
		err := client.HMSet(context.Background(), key, data).Err()
		if err != nil {
			ServerError(w)
			return
		}
		err = client.Expire(context.Background(), key, time.Second*180).Err()
		if err != nil {
			ServerError(w)
			return
		}
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "otp is sent",
		})
		WriteJson(w, jsonData, 200)
	})
	return login
}
