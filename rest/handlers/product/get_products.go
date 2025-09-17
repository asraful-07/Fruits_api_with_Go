package product

import (
	"encoding/json"
	"net/http"
)

// Get all fruits
func (h *Handler) GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// repo থেকে list আনা
	fruits, err := h.fruitsRepo.List()
	if err != nil {
		http.Error(w, "Failed to fetch fruits list", http.StatusInternalServerError)
		return
	}

	// JSON encode করে পাঠানো
	if err := json.NewEncoder(w).Encode(fruits); err != nil {
		http.Error(w, "Failed to encode fruits list", http.StatusInternalServerError)
		return
	}
}
