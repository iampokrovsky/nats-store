package main

import (
	"github.com/pokrovsky-io/nats-store/config"
	"github.com/pokrovsky-io/nats-store/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
