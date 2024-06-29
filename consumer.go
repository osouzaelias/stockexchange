package main

import "fmt"

type MyConsumer struct{}

func (this MyConsumer) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		event := ringBuffer[lower&BufferMask]
		processOrder(event)
	}
}

func processOrder(event TradeEvent) {
	fmt.Printf("Processing TradeEvent: OrderID=%d, Price=%d, Quantity=%d, BuySell=%s\n",
		event.OrderID, event.Price, event.Quantity, event.BuySell)

	if event.BuySell == "Buy" {
		processBuyOrder(event)
	} else {
		processSellOrder(event)
	}
}

func processBuyOrder(order TradeEvent) {
	for i := 0; i < len(orderBook.SellOrders); i++ {
		sellOrder := orderBook.SellOrders[i]
		if sellOrder.Price <= order.Price && sellOrder.Quantity >= order.Quantity {
			fmt.Printf("Executing Buy Order: OrderID=%d, Price=%d, Quantity=%d\n",
				order.OrderID, order.Price, order.Quantity)
			orderBook.SellOrders = append(orderBook.SellOrders[:i], orderBook.SellOrders[i+1:]...)
			return
		}
	}
	orderBook.BuyOrders = append(orderBook.BuyOrders, order)
}

func processSellOrder(order TradeEvent) {
	for i := 0; i < len(orderBook.BuyOrders); i++ {
		buyOrder := orderBook.BuyOrders[i]
		if buyOrder.Price >= order.Price && buyOrder.Quantity >= order.Quantity {
			fmt.Printf("Executing Sell Order: OrderID=%d, Price=%d, Quantity=%d\n",
				order.OrderID, order.Price, order.Quantity)
			orderBook.BuyOrders = append(orderBook.BuyOrders[:i], orderBook.BuyOrders[i+1:]...)
			return
		}
	}
	orderBook.SellOrders = append(orderBook.SellOrders, order)
}

type OrderBook struct {
	BuyOrders  []TradeEvent
	SellOrders []TradeEvent
}

var orderBook = OrderBook{
	BuyOrders:  make([]TradeEvent, 0),
	SellOrders: make([]TradeEvent, 0),
}
