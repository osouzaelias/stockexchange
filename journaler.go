package main

import "log"

type Journaler struct{}

func (j Journaler) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		event := ringBuffer[lower&BufferMask]
		// Lógica para registrar o evento
		log.Printf("Journaling TradeEvent: %+v\n", event)
		// Implementar lógica de persistência
	}
}
