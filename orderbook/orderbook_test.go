package orderbook

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrderbook(t *testing.T) {
	Convey("Test CreateOrderBook", t, func() {
		orderbook := CreateOrderBook(xrpjpy, sell)
		So(orderbook.Pair, ShouldEqual, xrpjpy)
		So(orderbook.LimitedOrders.Size(), ShouldEqual, 0)

		for i := 1; i <= 100; i++ {
			o := CreateOrder(limited, buy, xrpjpy, 0.62, float64(i))
			orderbook.AddOrder(o)
		}

		So(orderbook.LimitedOrders.Size(), ShouldEqual, 1)
	})

	Convey("Test AllLimitedOrders", t, func() {
		orderbook := CreateOrderBook(xrpjpy, sell)
		for i := 1; i <= 100; i++ {
			o := CreateOrder(limited, buy, xrpjpy, 0.61, float64(i))
			orderbook.AddOrder(o)
		}
		orders := orderbook.AllLimitedOrders()
		So(len(orders), ShouldEqual, 100)
	})
}
