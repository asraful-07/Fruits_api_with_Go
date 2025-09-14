package user

import (
	"encoding/json"
	"fruits-api/config"
	"fruits-api/global_product"
	"fruits-api/utils"
	"net/http"
)

type ReqUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login ReqUser 

	_= json.NewDecoder(r.Body).Decode(&login)

	usr := global_product.Find(login.Email, login.Password)

	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.JwtSecretKey, utils.Payload{
		Sub: usr.ID,
		FullName: usr.FullName,
		Email: usr.Email,
	})

	if  err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  usr,
		"token": accessToken,
	})
}