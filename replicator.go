package main

import "log"

type Replicator struct{}

func (r Replicator) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		event := ringBuffer[lower&BufferMask]
		// Lógica para replicar o evento para outros sistemas ou nós
		log.Printf("Replicating TradeEvent: %+v\n", event)
	}
}
