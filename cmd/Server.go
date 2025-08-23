package cmd

import (
	"fmt"
	"fruits-api/database"
	"fruits-api/middleware"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	// Manager বানানো
	manager := middleware.NewManager()

	// Global Middleware
	manager.Use(
		middleware.Logger,
		middleware.CORS,
		middleware.Preflight,
	)

	// Sample data load
	database.InitFruits()

	// Route registration আলাদা ফাইল থেকে
	RegisterRoutes(mux, manager)

	// Middleware wrap
	wrappedMux := manager.WrapMux(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
