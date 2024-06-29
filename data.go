package main

type (
	TradeEvent struct {
		OrderID  int64  `json:"order_id"`
		Price    int64  `json:"price"`
		Quantity int64  `json:"quantity"`
		BuySell  string `json:"buy_sell"`
	}

	OrderBook struct {
		BuyOrders  []TradeEvent
		SellOrders []TradeEvent
	}
)

const (
	BufferSize   = 1024 * 64
	BufferMask   = BufferSize - 1
	Reservations = 1
)

var ringBuffer [BufferSize]TradeEvent

var orderBook = OrderBook{
	BuyOrders:  make([]TradeEvent, 0),
	SellOrders: make([]TradeEvent, 0),
}
