package orderbook

import "github.com/davecgh/go-spew/spew"

func main() {
	// Create orderbook
	orderbook := CreateOrderBook(bchjpy, buy)

	// Add buy order to orderbook
	buyOrder := CreateOrder(limited, buy, xrpjpy, 100.1, 100)
	orderbook.AddOrder(buyOrder)

	// Add sell order to orderbook
	sellOrder := CreateOrder(limited, sell, xrpjpy, 110.1, 100)
	orderbook.AddOrder(sellOrder)

	// Delete order from orderbook
	//orderbook.DeleteOrder(sellOrder)

	// Strike the orderbook

	spew.Dump(orderbook)
}
