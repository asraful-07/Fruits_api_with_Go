package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
	JwtSecretKey string
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

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		// first time
	loadConfig()
	}
	return configuration
}
