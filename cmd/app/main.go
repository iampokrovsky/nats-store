package main

import (
	"nats-store/config"
	"nats-store/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
