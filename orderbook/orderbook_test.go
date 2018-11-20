package orderbook

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// STRIKE ORDER ==============================================
// when orderbook is thick enough
// when orderbook is not thick enough

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

	Convey("Given buy orderbook", t, func() {
		orderbook := prepareOrderbook(buy, 0.61)

		Convey("should able to DeleteOrder", func() {
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
		})

		Convey("When strike by a buy order", func() {
			myOrder := CreateOrder(limited, buy, xrpjpy, 0.68, 88)

			Convey("should return nil", func() {
				trades := orderbook.Strike(myOrder)
				So(trades, ShouldEqual, nil)
			})
		})

		Convey("When strike by a sell order", func() {
			myOrder := CreateOrder(limited, sell, xrpjpy, 0.68, 88)
			Convey("should be some trades which has higher price", func() {
				trades := orderbook.Strike(myOrder)
				for _, trade := range *trades {
					So(trade.Price, ShouldBeGreaterThanOrEqualTo, 0.68)
				}
				So(len(*trades), ShouldBeGreaterThan, 0)
			})
		})
	})

	Convey("Given sell orderbook", t, func() {
		orderbook := prepareOrderbook(sell, 0.61)
		Convey("When strike by a sell order", func() {
			myOrder := CreateOrder(limited, sell, xrpjpy, 0.68, 88)

			Convey("should return nil", func() {
				trades := orderbook.Strike(myOrder)
				So(trades, ShouldEqual, nil)
			})
		})

		Convey("When strike by a buy order", func() {
			myOrder := CreateOrder(limited, buy, xrpjpy, 0.68, 88)
			Convey("should be some trades which has lower price", func() {
				trades := orderbook.Strike(myOrder)
				for _, trade := range *trades {
					So(trade.Price, ShouldBeLessThanOrEqualTo, 0.68)
				}
				So(len(*trades), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func prepareOrderbook(side Side, price float64) *OrderBook {
	orderbook := CreateOrderBook(xrpjpy, side)
	for i := 0; i < 100; i++ {
		p := price * (100 + float64(i)) / 100
		order := CreateOrder(limited, side, xrpjpy, p, 10)
		orderbook.AddOrder(order)
	}

	return orderbook
}
