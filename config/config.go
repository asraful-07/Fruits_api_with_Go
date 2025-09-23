package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	if err:= godotenv.Load(); err != nil {
		fmt.Println("Failed to load the env variables")
		os.Exit(1)
	}
	
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPortStr, 10, 64)
	if err != nil {
		fmt.Println("Http port must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("jwt is required")
	    os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("dbPort is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 54) 
	if err != nil {
		fmt.Println("dbPort must be a number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("dbName is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("dbUser is required")
		os.Exit(1)
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("dbPassword is required")
		os.Exit(1)
	}

	dbSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enblSSLMode, err := strconv.ParseBool(dbSslMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host: dbHost,
		Port: int(dbPrt),
		Name: dbName,
		User: dbUser,
		Password: dbPass,
		EnableSSLMODE: enblSSLMode,
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		JwtSecretKey: jwtSecretKey,
		DB: dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
	// first time lazy loading
	loadConfig()
	}
	return configuration
}
