package rest

import (
	"fmt"
	"fruits-api/config"
	"fruits-api/database"
	"fruits-api/rest/middleware"
	"log"
	"net/http"
	"strconv"
)

func Start(cfg config.Config) {
	mux := http.NewServeMux()

	// Manager বানানো
	manager := middleware.NewManager()

	// Global Middleware
	manager.Use(
		middleware.Preflight,
		middleware.CORS,
		middleware.Logger,
	)

	// Sample data load
	database.InitFruits()

	// Route registration আলাদা ফাইল থেকে
	RegisterRoutes(mux, manager)

	// Middleware wrap
	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(int(cfg.HttpPort)) //type casting (inter to string)

	fmt.Println("Server running on http://localhost:" + addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}