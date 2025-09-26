package user

import (
	"fruits-api/utils"
	"net/http"
)

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    
	users, err := h.svc.List()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch fruits list")
		return
	}

	utils.SendData(w, users, http.StatusOK)
}