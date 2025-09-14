package cmd

import (
	"fruits-api/config"
	"fruits-api/rest"
	"fruits-api/rest/handlers/product"
	"fruits-api/rest/handlers/user"
)

func Serve() {
	cfg := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler)

    server.Start(cfg)
}
