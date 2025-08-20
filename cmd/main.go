package main

import (
	"fmt"
	"fruits-api/database"
	"fruits-api/handlers"
	"fruits-api/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Manager বানানো
	manager := middleware.NewManager()

	// Global Middleware (সব রিকোয়েস্টে চলবে)
	manager.Use(
		middleware.Logger,
		middleware.Test,
		middleware.CORS, 
	)

	// Sample data load
	database.InitFruits()

	// Routes
	mux.Handle("GET /fruits", manager.With(http.HandlerFunc(handlers.GetFruits)))
	mux.Handle("POST /fruits", manager.With(http.HandlerFunc(handlers.CreateFruits)))
	mux.Handle("GET /fruits/{id}", manager.With(http.HandlerFunc(handlers.GetById)))
	mux.Handle("PUT /fruits-update/{id}", manager.With(http.HandlerFunc(handlers.GetByUpdate)))
	mux.Handle("DELETE /fruits-delete/{id}", manager.With(http.HandlerFunc(handlers.GetByDelete)))

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
