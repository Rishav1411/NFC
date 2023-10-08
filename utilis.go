package main

import (
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "route does not exist",
	})
	w.Write(jsonData)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": "method is not valid",
	})
	w.Write(jsonData)
}
