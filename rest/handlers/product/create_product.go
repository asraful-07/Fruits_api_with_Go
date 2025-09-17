package product

import (
	"encoding/json"
	"fruits-api/repo"
	"net/http"
)

// Request body struct
type ReqCreateFruits struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}


// Fruits Create API
func (h *Handler) CreateFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ReqCreateFruits
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Repo struct বানানো
	newFruit := repo.Fruits{
		Name:        req.Name,
		Color:       req.Color,
		Image:       req.Image,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Description: req.Description,
	}

	// Repo এর Create method কল
	createdFruit, err := h.fruitsRepo.Create(newFruit)
	if err != nil {
		http.Error(w, "Failed to create fruit", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdFruit)
}
