package talib4g

import (
	"github.com/shopspring/decimal"
)

type OrderType int

const (
	BUY OrderType = iota
	SELL
)

type Order struct {
	Type   OrderType
	Price  decimal.Decimal
	Amount decimal.Decimal
}

func NewOrder(orderType OrderType) (o *Order) {
	o = new(Order)
	o.Type = orderType

	return o
}
