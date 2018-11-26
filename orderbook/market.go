package orderbook

// Market should have bid and ask orders
type Market struct {
	Pair          Pair // XRPJPY
	SellOrderBook *OrderBook
	BuyOrderBook  *OrderBook
}

// Strike the orderbook to check if any matching orders and generator trades
//
// Condition 1:
// order => buy, ¥10.1
// orderbook => sell 10.5, 11, 11.1
// result => go to Buy orderbook
//
// Condition 2:
// order => buy, ¥10.1
// orderbook => sell 9.5, 10, 10.2
// result => go to Buy orderbook
//
func (market *Market) Strike(myOrder *Order) *[]*Trade {
	var orderbook *OrderBook
	var counterOrderbook *OrderBook
	trades := &[]*Trade{}

	switch myOrder.Side {
	case buy:
		orderbook = market.SellOrderBook
		counterOrderbook = market.BuyOrderBook
	case sell:
		orderbook = market.BuyOrderBook
		counterOrderbook = market.SellOrderBook
	}

	if myOrder.Side == buy {
		// buy order is scanning selling part from low -> high.
		iter := orderbook.LimitedOrders.Iterator()
		iter.Begin()
		// loop the pricelevel
		for iter.Next() {
			priceLevel, _ := iter.Value().(*PriceLevel)
			// loop the orders in pricelevel
			for _, order := range *priceLevel.Orders {
				// 继续执行条件:
				// - 有符合价格的订单
				// - 订单没有被fill
				if myOrder.Price >= priceLevel.Price && myOrder.Amount > 0 {
					if myOrder.Amount > order.Amount {
						*trades = append(*trades, &Trade{Price: order.Price, Amount: order.Amount})
						myOrder.Amount = myOrder.Amount - order.Amount
						order.Amount = 0
						continue
					} else if myOrder.Amount < order.Amount {
						*trades = append(*trades, &Trade{Price: order.Price, Amount: myOrder.Amount})
						order.Amount = order.Amount - myOrder.Amount
						myOrder.Amount = 0
						break
					}
				} else if myOrder.Amount == 0 {
					break
				}
			}

			// striked order should be deleted
			var orders []*Order
			for _, order := range *priceLevel.Orders {
				if order.Amount != 0 {
					orders = append(orders, order)
				}
			}
			*priceLevel.Orders = orders
		}
	} else if myOrder.Side == sell {
		// sell order is scanning buying part from high -> low.
		iter := orderbook.LimitedOrders.Iterator()
		iter.End()
		for iter.Prev() {
			priceLevel, _ := iter.Value().(*PriceLevel)
			// loop the orders in pricelevel
			for _, order := range *priceLevel.Orders {
				// 继续执行条件:
				// - 有符合价格的订单
				// - 订单没有被fill
				if myOrder.Price <= priceLevel.Price && myOrder.Amount > 0 {
					if myOrder.Amount > order.Amount {
						// 不能 Fill 订单
						*trades = append(*trades, &Trade{Price: order.Price, Amount: order.Amount})
						myOrder.Amount = myOrder.Amount - order.Amount
						order.Amount = 0
						continue
					} else if myOrder.Amount < order.Amount {
						// 可以 Fill 订单
						*trades = append(*trades, &Trade{Price: order.Price, Amount: myOrder.Amount})
						order.Amount = order.Amount - myOrder.Amount
						myOrder.Amount = 0
						break
					}
				} else if myOrder.Amount == 0 {
					break
				}
			}

			// striked order should be deleted
			var orders []*Order
			for _, order := range *priceLevel.Orders {
				if order.Amount != 0 {
					orders = append(orders, order)
				}
			}
			*priceLevel.Orders = orders
		}
	}

	// 退出条件:
	// - 订单被fill
	// - 订单没被fill，没有符合价格的订单了
	// - 订单没被fill，没有订单了
	if myOrder.Amount > 0 {
		// but should be opposite side orderbook
		counterOrderbook.AddOrder(myOrder)
	}

	return trades
}

// CreateMarket will create a market object
func CreateMarket(pair Pair) *Market {
	return &Market{
		Pair:          pair,
		SellOrderBook: CreateOrderBook(pair, sell),
		BuyOrderBook:  CreateOrderBook(pair, buy),
	}
}
