package user

import (
	"fruits-api/rest/middleware"
	"net/http"
)

// RegisterRoutes all route
func (h *Handler) RegisterUserRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login",  manager.With(http.HandlerFunc(h.Login)))
	mux.Handle("GET /users",  manager.With(http.HandlerFunc(h.GetLogin)))
	mux.Handle("DELETE /users", manager.With(http.HandlerFunc(h.DeleteUsers)))
}
