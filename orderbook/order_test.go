package orderbook

import "testing"

func TestCreateOrder(t *testing.T) {
	order := CreateOrder(limited, buy, xrpjpy, 0.6, 10000)
	if order.Amount != 10000 || order.Pair != xrpjpy || order.Side != buy {
		t.Error("Order is not correct!")
	}
}
