package orderbook

// Trade struct
type Trade struct {
	ID         uint
	SequenceNo uint
	BuyerID    uint
	SellID     uint
	Pair       Pair
	Price      float64
	Amount     float64
	TotalPrice float64
	AskOrderID uint
	BidOrderID uint
}

func (trade *Trade) calculateTotalPrice() {
	trade.TotalPrice = trade.Price * trade.Amount
}

// CreateTrade : for creating a trade records
func CreateTrade(price float64, amount float64, pair Pair) {
	trade := &Trade{}
	trade.Price = price
	trade.Amount = amount
	trade.Pair = pair
	trade.calculateTotalPrice()
}
