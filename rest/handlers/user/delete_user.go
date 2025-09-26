package user

import (
	"fruits-api/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	if idStr == "" {
		utils.SendError(w, http.StatusBadRequest, "Missing id path")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Missing id path")
		return
	}

	 err = h.svc.Delete(id)
	 if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete fruit")
		return
	 }

	 utils.SendData(w, "User deleted successfully", http.StatusOK)
}