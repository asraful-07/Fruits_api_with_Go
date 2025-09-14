package product

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func (h *Handler) GetByUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	for index, fruit := range global_product.FruitsList {
		if fruit.ID == id {
			var updatedFruit global_product.Fruits
			if err := json.NewDecoder(r.Body).Decode(&updatedFruit); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			updatedFruit.ID = id
			global_product.FruitsList[index] = updatedFruit

			json.NewEncoder(w).Encode(updatedFruit)
			return
		}
	}

	http.Error(w, "Fruit not found", http.StatusNotFound)
}