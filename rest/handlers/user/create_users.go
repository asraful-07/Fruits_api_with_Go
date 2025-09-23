package user

import (
	"encoding/json"
	"fruits-api/domain"
	"fruits-api/utils"
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
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Map request data to domain.User
	newUser := domain.User{
		FullName:     req.FullName,
		Email:        req.Email,
		Password:     req.Password,  
		IsShopeOwner: req.IsShopeOwner,
	}

	// Call service layer
	createdUser, err := h.svc.Create(newUser)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendData(w, createdUser, http.StatusCreated)
}
