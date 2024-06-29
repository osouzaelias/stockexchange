package main

import (
	"encoding/json"
	"github.com/smartystreets-prototypes/go-disruptor"
	"log"
	"net/http"
)

func main() {
	myDisruptor := disruptor.New(
		disruptor.WithCapacity(BufferSize),
		disruptor.WithConsumerGroup(MyConsumer{}))

	go func() {
		http.HandleFunc("POST /order", func(w http.ResponseWriter, r *http.Request) {
			var order TradeEvent
			if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			publishOrder(myDisruptor, order)
			w.WriteHeader(http.StatusAccepted)
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	myDisruptor.Read()
}
