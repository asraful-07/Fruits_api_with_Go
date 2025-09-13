package product

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func (h *Handler) GetByDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	for index, fruit := range global_product.FruitsList {
		if fruit.ID == id {
			global_product.FruitsList = append(global_product.FruitsList[:index], global_product.FruitsList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Fruit Delete")
}