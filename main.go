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

	// Sample data load
	database.InitFruits()

	// Middleware যুক্ত করে route handle
	mux.Handle("GET /fruits", middleware.Logger(http.HandlerFunc(handlers.GetFruits)))
	mux.Handle("POST /fruits", middleware.Logger(http.HandlerFunc(handlers.CreateFruits)))
	mux.Handle("GET /fruits/{id}", middleware.Logger(http.HandlerFunc(handlers.GetById)))
	mux.Handle("PUT /fruits-update/{id}", middleware.Logger(http.HandlerFunc(handlers.GetByUpdate)))
	mux.Handle("DELETE /fruits-delete/{id}", middleware.Logger(http.HandlerFunc(handlers.GetByDelete)))

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},  // http://localhost:5173
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// 2025/08/16 18:54:35 [Logger] Completed in 503.1µs