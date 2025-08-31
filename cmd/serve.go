package cmd

import (
	"fruits-api/config"
	"fruits-api/rest"
)

func Serve() {
	cfg := config.GetConfig()
    rest.Start(cfg)
}
