package orderbook

import (
	"testing"
)

func TestCreateOrderbook(t *testing.T) {
	orderbook := CreateOrderBook(xrpjpy)
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
	orderbook := CreateOrderBook(xrpjpy)
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
	orderbook := CreateOrderBook(xrpjpy)
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
