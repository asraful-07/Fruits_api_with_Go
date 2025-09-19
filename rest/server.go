package rest

import (
	"fmt"
	"fruits-api/config"
	"fruits-api/rest/handlers/product"
	"fruits-api/rest/handlers/user"
	"fruits-api/rest/middleware"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cfg *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer(cfg *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return  &Server{
		cfg: cfg,
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()

	// Manager বানানো
	manager := middleware.NewManager()

	// Global Middleware
	manager.Use(
		middleware.Preflight,
		middleware.CORS,
		middleware.Logger,
	)

	// Route registration আলাদা ফাইল থেকে
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterUserRoutes(mux, manager)

	// Middleware wrap
	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(int(server.cfg.HttpPort)) //type casting (inter to string)

	fmt.Println("Server running on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}