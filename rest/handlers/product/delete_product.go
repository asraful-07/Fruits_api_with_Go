package product

import (
	"fruits-api/utils"
	"net/http"
	"strconv"
)

// Fruits Delete API
func (h *Handler) GetByDelete(w http.ResponseWriter, r *http.Request) {
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
		utils.SendError(w, http.StatusBadRequest,"Invalid id")
		return
	}

	// repo call
	err = h.svc.Delete(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError,"Failed to delete fruit")
		return
	}

    utils.SendData(w, "Fruit deleted successfully", http.StatusOK)
}
