package main

// OrderBook is for storing all the pending orders.
type OrderBook struct {
	Pair          Pair // XRPJPY
	LimitedOrders []PriceLevel
	MarketOrders  []PriceLevel
}

// func (orderbook *OrderBook) Match(order Order) {
// }

// func (orderbook *OrderBook) DeleteOrder(order Order) {
// }

// AddOrder add order to price level
func (orderbook *OrderBook) AddOrder(order Order) {
	pricelevel := PriceLevel{Price: order.Price, Volume: order.Amount, Orders: []Order{}}
	pricelevel.Orders = append(pricelevel.Orders, order)
	orderbook.LimitedOrders = append(orderbook.LimitedOrders, pricelevel)
}

// NewOrderBook : Create a new orderbook
func NewOrderBook() OrderBook {
	orderbook := OrderBook{Pair: xrpjpy, LimitedOrders: []PriceLevel{}}
	return orderbook
}
