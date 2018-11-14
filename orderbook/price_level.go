package orderbook

// PriceLevel in orderbook
type PriceLevel struct {
	Price  float64
	Volume float64
	Orders *[]*Order
}
