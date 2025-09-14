package product

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	for _, fruit := range global_product.FruitsList {
		if fruit.ID == id {
			json.NewEncoder(w).Encode(fruit)
			return
		}
	}
	http.Error(w, "Fruit not found", http.StatusNotFound)
}