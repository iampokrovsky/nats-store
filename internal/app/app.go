package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"github.com/pokrovsky-io/nats-store/config"
	"github.com/pokrovsky-io/nats-store/internal/repo"
	"github.com/pokrovsky-io/nats-store/internal/transport/nats"
	"github.com/pokrovsky-io/nats-store/internal/transport/rest"
	"github.com/pokrovsky-io/nats-store/internal/usecase"
	"github.com/pokrovsky-io/nats-store/pkg/httpserver"
	"github.com/pokrovsky-io/nats-store/pkg/postgres"
	"log"
	"sync"
)

func Run(cfg *config.Config) {

	// Repository
	pg, err := postgres.New(cfg.DB.GetURL())
	if err != nil {
		// TODO: Обработать ошибку
		log.Fatal(err)
	}
	defer pg.Close()

	r := repo.New(pg)

	// Use case
	uc := usecase.New(r)

	// STAN
	sc, err := stan.Connect(cfg.NATS.ClusterID, cfg.NATS.ClientID)
	if err != nil {
		// TODO: Обработать ошибку
		log.Fatal(err)
	}
	defer sc.Close()

	stn := nats.New(sc, uc)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go stn.Subscribe(wg, cfg.NATS.Subject)

	// HTTP Server
	handler := gin.New()
	rest.NewRouter(handler, uc)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// TODO: Обработать ошибку
	httpServer.Run()

	wg.Wait()

	fmt.Println("DONE")
}
