package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func validation(userData []byte) *User {
	user := &User{}
	json.Unmarshal(userData, user)
	validate := validator.New()
	validate.RegisterValidation("phone", PhoneValidor)
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
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "user is created",
		})
		WriteJson(w, jsonData, 201)
	})
	return sign_up
}
