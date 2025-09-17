package product

import (
	"encoding/json"
	"fruits-api/repo"
	"net/http"
	"strconv"
)

// Request body struct
type ReqUpdateFruits struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Update fruit by ID
func (h *Handler) GetByUpdate(w http.ResponseWriter, r *http.Request) {
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

	// Request body decode
	var req ReqUpdateFruits
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Repo struct বানানো
	updatedFruit := repo.Fruits{
		ID:          id,
		Name:        req.Name,
		Color:       req.Color,
		Image:       req.Image,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Description: req.Description,
	}

	// Repo update call
	fruit, err := h.fruitsRepo.Update(updatedFruit)
	if err != nil {
		http.Error(w, "Failed to update fruit", http.StatusInternalServerError)
		return
	}

	if fruit == nil {
		http.Error(w, "Fruit not found", http.StatusNotFound)
		return
	}

	// JSON encode response
	json.NewEncoder(w).Encode(fruit)
}
