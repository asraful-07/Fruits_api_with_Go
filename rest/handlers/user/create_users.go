package user

import (
	"encoding/json"
	"fruits-api/repo"
	"net/http"
)

// Request body struct
type ReqCreateUser struct {
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopeOwner bool   `json:"isShopeOwner"`
}

// Create user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ReqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := repo.User{
		FullName:     req.FullName,
		Email:        req.Email,
		Password:     req.Password,
		IsShopeOwner: req.IsShopeOwner,
	}

	createdUser, err := h.userRepo.Create(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
