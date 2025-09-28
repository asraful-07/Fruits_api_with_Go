package product

import (
	// "fruits-api/domain"
	"fruits-api/utils"
	"net/http"
	"strconv"
)

// Get all fruits
func (h *Handler) GetFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqQuery := r.URL.Query()

	pageAsStr  := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	
	// parse query params
	page, err := strconv.Atoi(pageAsStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitAsStr)
	if err != nil || limit <= 0 {
		limit = 4
	}

	// Service → Repo → DB clean architecture maintain 
	fruits, err := h.svc.List(int64(page), int64(limit)) 
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}

	utils.SendData(w, fruits, http.StatusOK)
}
