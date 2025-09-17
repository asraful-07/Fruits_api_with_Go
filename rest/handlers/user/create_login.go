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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Repo থেকে user খোঁজা
	usr, err := h.userRepo.Find(login.Email, login.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, utils.Payload{
		Sub:      usr.ID,
		FullName: usr.FullName,
		Email:    usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  usr,
		"token": accessToken,
	})
}
