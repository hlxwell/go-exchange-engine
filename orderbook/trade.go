package main

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
