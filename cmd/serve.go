package cmd

import (
	"fmt"
	"fruits-api/config"
	"fruits-api/fruits"
	"fruits-api/infra/db"
	"fruits-api/repo"
	"fruits-api/rest"
	prdHandler "fruits-api/rest/handlers/product"
	usrHandler "fruits-api/rest/handlers/user"
	"fruits-api/rest/middleware"
	"fruits-api/user"
	"os"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NewConnection(cfg.DB) 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	fruitsRepo := repo.NewFruitsRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	frtSvc := fruits.NewService(fruitsRepo)

    // middleware
	middlewares := middleware.NewMiddlewares(cfg)
	
	// handlers
	productHandler := prdHandler.NewHandler(middlewares, frtSvc)
	userHandler := usrHandler.NewHandler(cfg, usrSvc)
    
	// start sever
	server := rest.NewServer(cfg, userHandler, productHandler)
	server.Start()
}
