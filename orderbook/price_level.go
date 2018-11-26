package orderbook

// PriceLevel in orderbook
type PriceLevel struct {
	Price  float64
	Volume float64
	Orders *[]*Order
}

// HasBlankOrder will return if the price level has 0 volume order
func (priceLevel *PriceLevel) HasBlankOrder() bool {
	if len(*priceLevel.Orders) == 0 {
		return false
	}

	firstOrder := (*priceLevel.Orders)[0]
	if firstOrder.Amount == 0 {
		return true
	}

	return false
}
