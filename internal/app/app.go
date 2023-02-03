package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"nats-store/config"
	"nats-store/internal/repo"
	"nats-store/internal/transport/nats"
	"nats-store/internal/transport/rest"
	"nats-store/internal/usecase"
	"nats-store/pkg/httpserver"
	"nats-store/pkg/postgres"
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
