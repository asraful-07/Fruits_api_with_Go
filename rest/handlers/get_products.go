package handlers

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global_product.FruitsList)
}