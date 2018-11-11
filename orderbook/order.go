package orderbook

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
func CreateOrder(orderType OrderType, side Side, pair Pair, price float64, amount float64) *Order {
	order := &Order{Side: side, Pair: pair, Price: price, Amount: amount}
	order.ID = 1         // should be filled by DB
	order.SequenceNo = 1 // sequencer should fill it.
	order.CreatedAt = time.Now()
	order.Status = active
	order.Type = orderType
	order.OriginFunds = order.Price * order.Amount
	order.LeftFunds = order.OriginFunds
	return order
}
