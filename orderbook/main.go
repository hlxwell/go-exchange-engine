package main

import "github.com/davecgh/go-spew/spew"

func main() {
	orderbook := OrderBook{
		Pair:          bchjpy,
		LimitedOrders: []PriceLevel{},
		MarketOrders:  []PriceLevel{},
	}

	order := CreateOrder()
	orderbook.AddOrder(order)

	spew.Dump(orderbook)
}
