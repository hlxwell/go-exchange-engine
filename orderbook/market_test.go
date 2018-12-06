package orderbook

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMarket(t *testing.T) {
	Convey("CreateMarket", t, func() {
		market := CreateMarket(xrpjpy)
		market.BuyOrderBook = prepareOrderbook(buy, 0.61)
		market.SellOrderBook = prepareOrderbook(sell, 0.61)

		Convey("should create correct market", func() {
			So(market.Pair, ShouldEqual, xrpjpy)
			So(market.BuyOrderBook.Side, ShouldEqual, buy)
			So(market.SellOrderBook.Side, ShouldEqual, sell)
		})

		// Convey("should not be able to add Buy order to SellOrderbook", func() {
		// })

		Convey("should able to DeleteOrder", func() {
			order := CreateOrder(limited, buy, xrpjpy, 0.61, 100)
			market.BuyOrderBook.AddOrder(order)

			deleted, _ := market.BuyOrderBook.DeleteOrder(order)
			So(deleted, ShouldBeTrue)

			orders := market.BuyOrderBook.AllLimitedOrders()
			So(len(orders), ShouldBeGreaterThan, 0)
		})

		// when orderbook is thick enough

		Convey("When strike by a sell order", func() {
			myOrder := CreateOrder(limited, sell, xrpjpy, 0.68, 88)
			Convey("should be some trades which has higher price", func() {
				trades := market.Strike(myOrder)
				for _, trade := range *trades {
					So(trade.Price, ShouldBeGreaterThanOrEqualTo, 0.68)
				}
				So(len(*trades), ShouldBeGreaterThan, 0)
			})
		})

		Convey("When strike by a buy order", func() {
			myOrder := CreateOrder(limited, buy, xrpjpy, 0.68, 88)
			Convey("should be some trades which has lower price", func() {
				trades := market.Strike(myOrder)
				for _, trade := range *trades {
					So(trade.Price, ShouldBeLessThanOrEqualTo, 0.68)
				}
				So(len(*trades), ShouldBeGreaterThan, 0)
			})
		})

		// when orderbook is not thick enough

		Convey("When an sell order cannot be fullfilled", func() {
			myOrder := CreateOrder(limited, sell, xrpjpy, 0.61, 1200)
			market.Strike(myOrder)

			Convey("should have some left for this order", func() {
				So(myOrder.Amount, ShouldEqual, 200)
			})

			Convey("should add order to counter orderbook", func() {
				So(market.SellOrderBook.AllLimitedOrders(), ShouldContain, myOrder)
			})

			Convey("should delete all striked orders", func() {
				So(len(market.BuyOrderBook.AllLimitedOrders()), ShouldBeZeroValue)
			})
		})

		Convey("When an buy order cannot be fullfilled", func() {
			myOrder := CreateOrder(limited, buy, xrpjpy, 1.5, 1200)
			market.Strike(myOrder)

			Convey("should have some left for this order", func() {
				So(myOrder.Amount, ShouldEqual, 200)
			})

			Convey("should add order to counter orderbook", func() {
				So(market.BuyOrderBook.AllLimitedOrders(), ShouldContain, myOrder)
			})

			Convey("should delete all striked orders", func() {
				So(len(market.SellOrderBook.AllLimitedOrders()), ShouldBeZeroValue)
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
