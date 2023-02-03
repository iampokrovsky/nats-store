package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"nats-store/internal/entity"
	"nats-store/internal/usecase"
	"sync"
)

type STAN struct {
	conn    stan.Conn
	useCase usecase.Order
}

func New(sc stan.Conn, uc usecase.Order) *STAN {
	return &STAN{
		conn:    sc,
		useCase: uc,
	}
}

// TODO  Переименовать функцию
func (stn *STAN) inputHandler(msg *stan.Msg) {
	var order entity.Order
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		// TODO обработать ошибку
		log.Fatal(err)
	}

	// TODO: Накапливать входящие результаты и коммитить их с опредленной частотой
	stn.useCase.Create(&order)
}

func (stn *STAN) Subscribe(wg *sync.WaitGroup, subj string) {
	// Use a WaitGroup to wait for a message to arrive

	// Subscribe
	_, err := stn.conn.Subscribe(subj, func(msg *stan.Msg) {
		//defer wg.Done()
		stn.inputHandler(msg)
	})
	if err != nil {
		// TODO: Обработать ошибку
		log.Fatal(err)
	}
}
