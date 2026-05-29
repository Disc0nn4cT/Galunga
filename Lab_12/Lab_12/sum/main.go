package main

import (
	"encoding/json"
	"log"
	"os"

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

	sum := 0
	done := make(chan bool)

	nc.Subscribe("pipeline.squared", func(m *nats.Msg) {
		var d Data
		json.Unmarshal(m.Data, &d)

		if d.Value == -1 {
			log.Printf("==================================================")
			log.Printf("Результат: Сума квадратів парних чисел = %d", sum)
			log.Printf("==================================================")
			done <- true
			return
		}

		sum += d.Value
	})

	log.Println("Sum service готовий...")
	<-done
	os.Exit(0)
}
