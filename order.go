package talib4g

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)

type Order struct {
	Type          OrderSide
	Price         decimal.Decimal
	Amount        decimal.Decimal
	ExecutionTime time.Time
}

func NewOrder(orderType OrderSide) (o *Order) {
	o = new(Order)
	o.Type = orderType

	return o
}
