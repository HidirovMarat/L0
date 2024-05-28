package nats

import (
	"context"
	"encoding/json"
	"log"
	"sub/internal/storage/post"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func New(clusterID, clientID, natsURL string) *stan.Conn {
	// Подключение к NATS Streaming Server
	clientID = clientID + uuid.NewString()
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка при подключении к NATS Streaming Server: %v", err)
	}
	return &sc
}

func MessageHandlerFunc(cash map[string]post.Order, pg *post.Postgres) stan.MsgHandler {
	return func(m *stan.Msg) {
		var order post.Order
		err := json.Unmarshal(m.Data, &order)

		if err != nil {
			log.Print("Error to saveOrder", err)
			return
		}

		cash[order.Order_uid] = order
		//Log

		err = pg.CreateOrder(context.Background(), order)
		if err != nil {
			log.Print("")
			return
		}
	}
}


