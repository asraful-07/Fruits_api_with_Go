package main

import (
	"fmt"
	"fruits-api/cmd"
	"fruits-api/config"
)

func main() {
	cfg := config.GetConfig()
    fmt.Println("Service:", cfg.ServiceName, "Port:", cfg.HttpPort)
	cmd.Server()
}