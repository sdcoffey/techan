package talib4g

import (
	"time"

	"github.com/sdcoffey/big"
)

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)

type order struct {
	Type          OrderSide
	Price         big.Decimal
	Amount        big.Decimal
	ExecutionTime time.Time
}

func NewOrder(orderType OrderSide) (o *order) {
	o = new(order)
	o.Type = orderType

	return o
}
