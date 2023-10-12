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

func validation(userData []byte) *User {
	user := &User{}
	unerr := json.Unmarshal(userData, user)
	if unerr != nil {
		return nil
	}
	validate := validator.New()
	validate.RegisterValidation("phone", PhoneValidor)
	validate.RegisterValidation("reg", RegValidator)
	err := validate.Struct(user)
	if err != nil {
		return nil
	}
	return user
}
func SignUp() *chi.Mux {
	sign_up := chi.NewRouter()
	sign_up.Post("/", func(w http.ResponseWriter, r *http.Request) {
		requestData, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		user := validation(requestData)
		if user == nil {
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
		res, err := operations.CheckUser(user.Phone, db)
		if err != nil {
			ServerError(w)
			return
		}
		if res != -1 {
			jsonData, _ := json.Marshal(map[string]interface{}{
				"details": "user already exists",
			})
			WriteJson(w, jsonData, http.StatusConflict)
			return
		}
		client := database.RedisConnection()
		if client == nil {
			ServerError(w)
			return
		}
		defer client.Close()
		otp := GenerateOTP()
		key := GenerateKey(user.Phone)
		if !SendSMS(user.Phone, otp) {
			ServerError(w)
			return
		}
		data := map[string]interface{}{
			"type":  "sign_up",
			"otp":   otp,
			"name":  user.Name,
			"phone": user.Phone,
			"reg":   user.Reg,
		}
		err = client.HMSet(context.Background(), key, data).Err()
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
			"details": "user is created",
		})
		WriteJson(w, jsonData, 201)
	})
	return sign_up
}
