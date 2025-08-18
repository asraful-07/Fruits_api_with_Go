package main

import (
	"fmt"
	"fruits-api/database"
	"fruits-api/handlers"
	"fruits-api/middleware"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	// Manager বানানো
	manager := middleware.NewManager()

	// Global Middleware (সব রিকোয়েস্টে চলবে)
	manager.Use(middleware.Logger, middleware.Test)

	// Sample data load
	database.InitFruits()

	// Routes (middleware ছাড়াই আগে register করি)
	mux.Handle("GET /fruits", manager.With(http.HandlerFunc(handlers.GetFruits)))
	mux.Handle("POST /fruits",  manager.With(http.HandlerFunc(handlers.CreateFruits)))
	mux.Handle("GET /fruits/{id}",  manager.With(http.HandlerFunc(handlers.GetById)))
	mux.Handle("PUT /fruits-update/{id}",  manager.With(http.HandlerFunc(handlers.GetByUpdate)))
	mux.Handle("DELETE /fruits-delete/{id}",  manager.With(http.HandlerFunc(handlers.GetByDelete)))
	

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // http://localhost:5173
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
