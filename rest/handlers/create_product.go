package handlers

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func CreateFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var fruits global_product.Fruits
	if err := json.NewDecoder(r.Body).Decode(&fruits); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	global_product.FruitsList = append(global_product.FruitsList, fruits)
	json.NewEncoder(w).Encode(fruits)
}