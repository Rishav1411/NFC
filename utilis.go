package main

import (
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "route does not exist",
	})
	w.Write(jsonData)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(405)
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "method is not valid",
	})
	w.Write(jsonData)
}
