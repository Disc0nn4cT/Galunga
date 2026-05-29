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

	nc.Subscribe("pipeline.even", func(m *nats.Msg) {
		var d Data
		json.Unmarshal(m.Data, &d)

		if d.Value == -1 {
			nc.Publish("pipeline.squared", m.Data)
			return
		}

		squared := d.Value * d.Value
		log.Printf("Square: %d^2 = %d", d.Value, squared)

		newData := Data{Value: squared}
		bytes, _ := json.Marshal(newData)
		nc.Publish("pipeline.squared", bytes)
	})

	log.Println("Square service готовий...")
	select {}
}
