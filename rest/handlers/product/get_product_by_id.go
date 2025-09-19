package product

import (
	"fruits-api/utils"
	"net/http"
	"strconv"
)

// Get fruit by ID
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
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

	// repo থেকে fruit আনো
	fruit, err := h.fruitsRepo.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruit")
		return
	}

	if fruit == nil {
		utils.SendError(w, http.StatusNotFound, "Fruit not found")
		return
	}

	utils.SendData(w, fruit, http.StatusOK)
	
}
