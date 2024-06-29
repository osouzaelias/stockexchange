package main

type TradeEvent struct {
	OrderID  int64  `json:"order_id"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
	BuySell  string `json:"buy_sell"`
}

const (
	BufferSize   = 1024 * 64
	BufferMask   = BufferSize - 1
	Reservations = 1
)

var ringBuffer [BufferSize]TradeEvent
