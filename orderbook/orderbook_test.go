package orderbook

import (
	"testing"
)

func TestCreateOrderbook(t *testing.T) {
	orderbook := CreateOrderBook(xrpjpy, sell)
	if orderbook.Pair != xrpjpy {
		t.Error("Wrong pair code.")
	}

	if orderbook.LimitedOrders.Size() != 0 {
		t.Error("Should have no orders.")
	}

	for i := 1; i <= 100; i++ {
		o := CreateOrder(limited, buy, xrpjpy, 0.62, float64(i))
		orderbook.AddOrder(o)
	}

	if orderbook.LimitedOrders.Size() != 1 {
		t.Error("Expected limited order count should be 1.")
	}
}

func TestAllLimitedOrders(t *testing.T) {
	orderbook := CreateOrderBook(xrpjpy, sell)
	for i := 1; i <= 100; i++ {
		o := CreateOrder(limited, buy, xrpjpy, 0.61, float64(i))
		orderbook.AddOrder(o)
	}
	orders := orderbook.AllLimitedOrders()
	if len(orders) != 100 {
		t.Errorf("Expected limited order count should be 100, but it was %d", len(orders))
	}
}

func TestDeleteOrderbook(t *testing.T) {
	orderbook := CreateOrderBook(xrpjpy, sell)

	order := CreateOrder(limited, buy, xrpjpy, 0.61, 100)
	orderbook.AddOrder(order)

	deleted, _ := orderbook.DeleteOrder(order)
	if !deleted {
		t.Error("Expected delete the order from order book.")
	}

	orders := orderbook.AllLimitedOrders()
	if len(orders) != 0 {
		t.Errorf("Expected limited order count should be 0, but it was %d", len(orders))
	}
}

func TestStrike(t *testing.T) {
	orderbook := CreateOrderBook(xrpjpy, sell)
	for i := 0; i < 100; i++ {
		price := 0.61 * (100 + float64(i)) / 100
		order := CreateOrder(limited, buy, xrpjpy, price, 10)
		orderbook.AddOrder(order)
	}

	myOrder := CreateOrder(limited, buy, xrpjpy, 0.61, 100)
	orderbook.Strike(*myOrder)
}
