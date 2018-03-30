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

type Order struct {
	Type          OrderSide
	Security      string
	Price         big.Decimal
	Amount        big.Decimal
	ExecutionTime time.Time
	FeePercentage big.Decimal
}

func NewOrder(orderType OrderSide) (o *Order) {
	o = new(Order)
	o.Type = orderType

	return o
}
