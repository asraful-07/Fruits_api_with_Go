package product

import (
	"encoding/json"
	"fruits-api/domain"
	"fruits-api/utils"
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
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Repo struct বানানো
	newFruit := domain.Fruits{
		Name:        req.Name,
		Color:       req.Color,
		Image:       req.Image,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Description: req.Description,
	}

	// Repo এর Create method কল
	createdFruit, err := h.svc.Create(newFruit)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create fruit")
		return
	}

	utils.SendData(w, createdFruit, http.StatusCreated)
}
