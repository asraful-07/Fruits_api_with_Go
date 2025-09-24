package product

import (
	"fruits-api/utils"
	"net/http"
)

// Get all fruits
func (h *Handler) GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// repo থেকে list আনা
	fruits, err := h.svc.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}

	
	utils.SendData(w, fruits, http.StatusOK)
	
}
