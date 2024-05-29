package main

import (
	"encoding/json"
	"fmt"
	"pub/internal/app/nats"
	"pub/internal/entity"
	"time"

	"github.com/google/uuid"
)

func main() {

	clusterID := "test-cluster"
	clientID := uuid.NewString()
	natsURL := "nats://localhost:4222"

	// Подключение к NATS Streaming Server
	sc := nats.New(clusterID, clientID, natsURL)
	defer sc.Close()

	for i := 0; i < 20; i++ {

		jsonOrder := GetOrder()

		if i == 0 {
			jsonOrder.Order_uid = "text"
		}

		order, _ := json.Marshal(jsonOrder)

		err := sc.Publish("createOrder", order)
		fmt.Print(err)
	}
}

func GetOrder() entity.Order {
	dataBaseJSON :=
		`
	{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 9934930,
			"track_number": "WBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	  }
	`
	var orderBase entity.Order

	_ = json.Unmarshal([]byte(dataBaseJSON), &orderBase)

	orderBase.Order_uid = uuid.NewString()

	orderBase.Sm_id = 1000
	orderBase.Date_created = time.Now()

	return orderBase
}
