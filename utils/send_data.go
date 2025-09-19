package utils

import (
	"encoding/json"
	"net/http"
)

// Success response
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Error response
func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(msg);
}

// func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
// 	w.WriteHeader(StatusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(data)
// }

// func SendError(w http.ResponseWriter, StatusCode int, msg string) {
// 	w.WriteHeader(StatusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(msg)
// }