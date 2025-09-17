package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Get fruit by ID
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// URL থেকে id পড়া
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// string → int convert
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// repo থেকে fruit আনো
	fruit, err := h.fruitsRepo.Get(id)
	if err != nil {
		http.Error(w, "Failed to fetch fruit", http.StatusInternalServerError)
		return
	}

	if fruit == nil {
		http.Error(w, "Fruit not found", http.StatusNotFound)
		return
	}

	// JSON এ encode করে response
	if err := json.NewEncoder(w).Encode(fruit); err != nil {
		http.Error(w, "Failed to encode fruit", http.StatusInternalServerError)
		return
	}
}
