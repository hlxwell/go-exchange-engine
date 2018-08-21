package main

import "fmt"

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

// Order struct
type Order struct {
	ID int
}

// Trade struct
type Trade struct {
}

func main() {
	account := Account{Email: "hlxwell@gmail.com"}

	order := Order{ID: 1}

	fmt.Printf("hello shit %#v, %#v, %#v", account, order, XRP)
}
