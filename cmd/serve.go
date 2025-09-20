package cmd

import (
	"fruits-api/config"
	"fruits-api/repo"
	"fruits-api/rest"
	"fruits-api/rest/handlers/product"
	"fruits-api/rest/handlers/user"
	"fruits-api/rest/middleware"
)

func Serve() {
	cfg := config.GetConfig()

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(cfg, userRepo)

	middlewares := middleware.NewMiddlewares(cfg)

	fruitsRepo := repo.NewFruitsRepo()
	productHandler := product.NewHandler(middlewares, fruitsRepo)

	server := rest.NewServer(cfg, userHandler, productHandler)
   
	server.Start()
}
