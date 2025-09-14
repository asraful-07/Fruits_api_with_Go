package product

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func (h *Handler) GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(global_product.FruitsList); err != nil {
    http.Error(w, "Failed to encode fruits list", http.StatusInternalServerError)
    return
}
}