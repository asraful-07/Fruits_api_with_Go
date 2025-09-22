package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	// user => postgres
	// password => 123456789//Rahat@go
	// host => localhost
	// port => 5432
	// db name => ecommerce
	return  "user=postgres password=Rahat@go host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NecConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
 	}

	return  dbCon, nil
}