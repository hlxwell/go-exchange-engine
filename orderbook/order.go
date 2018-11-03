package main

import "time"

// Order data structure
type Order struct {
	ID          uint
	UserID      uint
	SequenceNo  uint
	Side        Side
	Pair        Pair
	Price       float64
	Amount      float64
	OriginFunds float64
	LeftFunds   float64
	Type        OrderType
	Status      OrderStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DoneAt      time.Time
}

// CreateOrder to create a new order
func CreateOrder() Order {
	order := Order{Side: buy, Pair: btcjpy, Price: 5000, Amount: 5}
	order.SequenceNo = 1
	order.CreatedAt = time.Now()
	order.Status = active
	order.Type = market
	order.OriginFunds = order.Price * order.Amount
	order.LeftFunds = order.OriginFunds
	return order
}
