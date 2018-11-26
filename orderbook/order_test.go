package orderbook

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrder(t *testing.T) {
	Convey("CreateOrder", t, func() {
		order := CreateOrder(limited, buy, xrpjpy, 0.6, 10000)

		Convey("Should have correct amount", func() {
			So(order.Amount, ShouldEqual, 10000)
		})

		Convey("Should have correct pair", func() {
			So(order.Pair, ShouldEqual, xrpjpy)
		})

		Convey("Should have correct side", func() {
			So(order.Side, ShouldEqual, buy)
		})
	})
}
