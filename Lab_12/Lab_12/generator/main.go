package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Data struct {
	Value int `json:"value"`
}

func main() {

	nc, err := nats.Connect("nats://nats-server:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	time.Sleep(2 * time.Second)
	log.Println("Generator started. Відправка чисел...")

	for i := 1; i <= 100; i++ {
		sendValue(nc, i)
		time.Sleep(5 * time.Millisecond)
	}

	sendValue(nc, -1)
	log.Println("Generator finished.")
}

func sendValue(nc *nats.Conn, val int) {
	data := Data{Value: val}
	bytes, _ := json.Marshal(data)
	nc.Publish("pipeline.numbers", bytes)
}
