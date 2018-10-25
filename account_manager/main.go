package main

// Currency enum type
type Currency int

// All the enum values of Currency type
const (
	JPY Currency = iota
	BTC
	XRP
	BCH
	ETH
)

// Account struct
type Account struct {
	Email             string
	AvailableBalances []Balance
	LockedBalances    []LockedBalance
}

// Balance struct
type Balance struct {
	CurrencyType Currency
	Amount       float32
}

// LockedBalance struct
type LockedBalance struct {
}
