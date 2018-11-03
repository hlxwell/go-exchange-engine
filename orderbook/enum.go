package main

// ENUM in GO https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3

// Pair : Asset Pair
type Pair int

const (
	btcjpy Pair = iota
	bchjpy
	xrpjpy
)

// OrderType : Order Type
type OrderType int

const (
	limited OrderType = iota
	market
)

// OrderStatus : Order Status
type OrderStatus int

const (
	active OrderStatus = iota
	done
)

// Side : Buy or Sell
type Side int

const (
	buy Side = iota
	sell
)
