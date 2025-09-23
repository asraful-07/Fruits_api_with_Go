package cmd

import (
	"fmt"
	"fruits-api/config"
	"fruits-api/infra/db"
	"fruits-api/repo"
	"fruits-api/rest"
	"fruits-api/rest/handlers/product"
	"fruits-api/rest/handlers/user"
	"fruits-api/rest/middleware"
	"os"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NecConnection(cfg.DB) 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cfg, userRepo)

	middlewares := middleware.NewMiddlewares(cfg)

	fruitsRepo := repo.NewFruitsRepo(dbCon)
	productHandler := product.NewHandler(middlewares, fruitsRepo)

	server := rest.NewServer(cfg, userHandler, productHandler)
   
	server.Start()
}
