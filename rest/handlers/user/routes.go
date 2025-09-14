package user

import (
	"fruits-api/rest/middleware"
	"net/http"
)

// RegisterRoutes সব route এখানে সেট হবে
func (h *Handler) RegisterUserRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login",  manager.With(http.HandlerFunc(h.Login)))
	mux.Handle("GET /users",  manager.With(http.HandlerFunc(h.GetLogin)))
}
