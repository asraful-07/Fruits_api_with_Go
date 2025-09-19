package product

import (
	"encoding/json"
	"fruits-api/repo"
	"fruits-api/utils"
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
		utils.SendError(w, http.StatusBadRequest, "Missing id parameter")
		return
	}

	// string → int convert
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// Request body decode
	var req ReqUpdateFruits
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid JSON")
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
		utils.SendError(w, http.StatusInternalServerError,"Failed to update fruit")
		return
	}

	if fruit == nil {
		utils.SendError(w,  http.StatusNotFound, "Fruit not found",)
		return
	}

	// JSON encode response
	utils.SendData(w, "Successfully updated data", http.StatusOK)
	
}
