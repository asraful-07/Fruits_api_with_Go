package user

import (
	"encoding/json"
	"fruits-api/utils"
	"net/http"
)

type ReqUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login ReqUser
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Repo Find To user 
	usr, err := h.svc.Find(login.Email, login.Password)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if usr == nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, utils.Payload{
		Sub:      usr.ID,
		FullName: usr.FullName,
		Email:    usr.Email,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, accessToken, http.StatusCreated)
}
