package cmd

import (
	"fruits-api/config"
	"fruits-api/rest"
	"fruits-api/rest/handlers/product"
	"fruits-api/rest/handlers/user"
	"fruits-api/rest/middleware"
)

func Serve() {
	cfg := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cfg)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(cfg, productHandler, userHandler)

    server.Start()
}
