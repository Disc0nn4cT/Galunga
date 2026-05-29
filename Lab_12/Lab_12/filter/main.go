package main

import (
	"encoding/json"
	"log"

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

	nc.Subscribe("pipeline.numbers", func(m *nats.Msg) {
		var d Data
		json.Unmarshal(m.Data, &d)

		if d.Value == -1 {
			nc.Publish("pipeline.even", m.Data)
			return
		}

		if d.Value%2 == 0 {
			log.Printf("Filter: пропущено парне число %d", d.Value)
			nc.Publish("pipeline.even", m.Data)
		}
	})

	log.Println("Filter service готовий...")
	select {} // Блокуємо завершення програми
}
