package main

import (
	"fmt"
)

type Pair string

// Order data structure
type Order struct {
	pair     Pair
	currency Currency
}

// Trade struct
type Trade struct {
}

// OrderBook is for storing all the pending orders.
type OrderBook struct {
	buyOrders  []Order
	sellOrders []Order
}

func init() {
	fmt.Println("initing orderbook")
}

func (orderbook *OrderBook) Match(order Order) {
}

func (orderbook *OrderBook) DeleteOrder(order Order) {
}

func (orderbook *OrderBook) AddOrder(order Order) {
}
