package rest

import (
	"fruits-api/rest/handlers"
	"fruits-api/rest/middleware"
	"net/http"
)

// RegisterRoutes সব route এখানে সেট হবে
func RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /fruits", manager.With(http.HandlerFunc(handlers.GetFruits)))
	mux.Handle("POST /fruits", manager.With(http.HandlerFunc(handlers.CreateFruits)))
	mux.Handle("GET /fruits/{id}", manager.With(http.HandlerFunc(handlers.GetById), middleware.Test))
	mux.Handle("PUT /fruits-update/{id}", manager.With(http.HandlerFunc(handlers.GetByUpdate)))
	mux.Handle("DELETE /fruits-delete/{id}", manager.With(http.HandlerFunc(handlers.GetByDelete)))
}
