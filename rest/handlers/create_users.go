package handlers

import (
	"encoding/json"
	"fruits-api/global_product"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user global_product.User 

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}
	global_product.UserList = append(global_product.UserList, user) 
	json.NewEncoder(w).Encode(http.StatusOK)
}