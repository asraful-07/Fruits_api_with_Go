package product

import (
	"fruits-api/rest/middleware"
	"net/http"
)

// RegisterRoutes সব route এখানে সেট হবে
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /fruits", manager.With(http.HandlerFunc(h.GetFruits)))
	mux.Handle("POST /fruits", manager.With(http.HandlerFunc(h.CreateFruits), h.middlewares.Authenticate))
	mux.Handle("GET /fruits/{id}", manager.With(http.HandlerFunc(h.GetById), middleware.Test))
	mux.Handle("PUT /fruits/{id}", manager.With(http.HandlerFunc(h.GetByUpdate), h.middlewares.Authenticate))
	mux.Handle("DELETE /fruits/{id}", manager.With(http.HandlerFunc(h.GetByDelete), h.middlewares.Authenticate))
}
