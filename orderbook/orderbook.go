package orderbook

import (
	"errors"
	"fmt"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// OrderBook is for storing all the pending orders.
type OrderBook struct {
	Pair          Pair // XRPJPY
	Side          Side // Buy or Sell
	LimitedOrders *rbt.Tree
	MarketOrders  *rbt.Tree
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
func (orderbook *OrderBook) Strike(order Order) bool {
	if order.Side == buy {
		// orderbook.LimitedOrders.Ceiling
		it := orderbook.LimitedOrders.Iterator()
		for it.Next() {
			priceLevel, _ := it.Value().(*PriceLevel)
			if priceLevel.Price < order.Price {
				fmt.Println(priceLevel)
			}
		}
	} else if order.Side == sell {
		// orderbook.LimitedOrders.Ceiling
	}

	return false
}

// GetPriceLevel get orders of a price level by price and order type
func (orderbook *OrderBook) GetPriceLevel(price float64, orderType OrderType) (*PriceLevel, error) {
	var iterator rbt.Iterator
	if orderType == limited {
		iterator = orderbook.LimitedOrders.Iterator()
	} else {
		iterator = orderbook.MarketOrders.Iterator()
	}

	for iterator.Next() {
		priceLevel := iterator.Value().(*PriceLevel)
		if priceLevel.Price == price {
			return priceLevel, nil
		}
	}

	return nil, errors.New("cannot find price level")
}

// DeleteOrder will delete order with ID.
func (orderbook *OrderBook) DeleteOrder(order *Order) (bool, error) {
	var (
		priceLevel *PriceLevel
		value      interface{}
		found      bool
	)

	if order.Type == limited {
		value, found = orderbook.LimitedOrders.Get(order.Price)
	} else {
		value, found = orderbook.MarketOrders.Get(order.Price)
	}

	if found {
		priceLevel = value.(*PriceLevel)
	}

	for i, o := range priceLevel.Orders {
		if o == order {
			priceLevel.Orders = append(priceLevel.Orders[:i], priceLevel.Orders[i+1:]...)
			return true, nil
		}
	}

	return false, errors.New("Cannot delete order")
}

// AddMarketOrder add market order
func (orderbook *OrderBook) AddMarketOrder(order *Order) {
	var priceLevel *PriceLevel
	value, found := orderbook.MarketOrders.Get(order.Price)
	if found {
		priceLevel = value.(*PriceLevel)
		priceLevel.Orders = append(priceLevel.Orders, order)
	} else {
		priceLevel = &PriceLevel{
			Price:  order.Price,
			Volume: order.Amount,
			Orders: []*Order{},
		}
		priceLevel.Orders = append(priceLevel.Orders, order)
	}
	orderbook.MarketOrders.Put(order.Price, priceLevel)
}

// AddLimitedOrder add limited order
func (orderbook *OrderBook) AddLimitedOrder(order *Order) {
	var priceLevel *PriceLevel
	value, found := orderbook.LimitedOrders.Get(order.Price)
	if found {
		priceLevel = value.(*PriceLevel)
		priceLevel.Orders = append(priceLevel.Orders, order)
	} else {
		priceLevel = &PriceLevel{
			Price:  order.Price,
			Volume: order.Amount,
			Orders: []*Order{},
		}
		priceLevel.Orders = append(priceLevel.Orders, order)
	}
	orderbook.LimitedOrders.Put(priceLevel.Price, priceLevel)
}

// AddOrder add order to price level
func (orderbook *OrderBook) AddOrder(order *Order) {
	switch order.Type {
	case market:
		orderbook.AddMarketOrder(order)
		break
	case limited:
		orderbook.AddLimitedOrder(order)
		break
	}
}

// AllLimitedOrders return all the limited orders
func (orderbook *OrderBook) AllLimitedOrders() []*Order {
	var orders []*Order
	i := orderbook.LimitedOrders.Iterator()
	for i.Next() {
		priceLevel := i.Value().(*PriceLevel)

		for _, order := range priceLevel.Orders {
			orders = append(orders, order)
		}
	}
	return orders
}

// CreateOrderBook : Create a new order book
func CreateOrderBook(pair Pair) *OrderBook {
	limitedOrders := rbt.NewWith(utils.Float64Comparator) // empty (keys are of type int)
	marketOrders := rbt.NewWith(utils.Float64Comparator)  // empty (keys are of type int)

	orderbook := &OrderBook{
		Pair:          pair,
		LimitedOrders: limitedOrders,
		MarketOrders:  marketOrders,
	}
	return orderbook
}
