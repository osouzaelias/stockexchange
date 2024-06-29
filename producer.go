package main

import "github.com/smartystreets-prototypes/go-disruptor"

func publishOrder(myDisruptor disruptor.Disruptor, order TradeEvent) {
	sequence := myDisruptor.Reserve(Reservations)

	ringBuffer[sequence&BufferMask] = order

	myDisruptor.Commit(sequence, sequence)
}
