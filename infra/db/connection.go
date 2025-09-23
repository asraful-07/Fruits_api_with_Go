package db

import (
	"fmt"
	"fruits-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cfg *config.DBConfig) string {
    connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s ", 
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

    if !cfg.EnableSSLMODE {
        connString += "sslmode=disable"
    }

    return connString
}

// return  "user=postgres password=Rahat@go host=localhost port=5432 dbname=ecommerce sslmode=disable"


func NewConnection(cfg *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cfg)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
 	}

	return  dbCon, nil
}

	// user => postgres
	// password => 123456789//Rahat@go
	// host => localhost
	// port => 5432
	// db name => ecommerce
	// sslmode=disable