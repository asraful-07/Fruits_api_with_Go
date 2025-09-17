package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Fruits Delete API
func (h *Handler) GetByDelete(w http.ResponseWriter, r *http.Request) {
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

	// repo call
	err = h.fruitsRepo.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete fruit", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Fruit deleted successfully",
	})
}
