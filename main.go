package main

import (
	"fmt"
)

func main() {
	order := Order{}
	orderbook := OrderBook{}
	orderbook.AddOrder(order)
	fmt.Println("order book.", orderbook)
}
