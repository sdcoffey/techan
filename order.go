package talib4g

import (
	"time"

	"github.com/sdcoffey/big"
)

// OrderSide is a simple enumeration representing the side of an Order (buy or sell)
type OrderSide int

// BUY and SELL enumerations
const (
	BUY OrderSide = iota
	SELL
)

// Order represents a trade execution (buy or sell) with associated metadata
type Order struct {
	Side          OrderSide
	Security      string
	Price         big.Decimal
	Amount        big.Decimal
	ExecutionTime time.Time
	FeePercentage big.Decimal
}

var one = big.NewFromString("1")

func (o Order) TotalAmount() big.Decimal {
	if o.Side == BUY {
		return o.Amount.Mul(one.Sub(o.FeePercentage))
	}

	return o.Amount
}

func (o Order) Profit() big.Decimal {
	profit := o.Amount.Mul(o.Price)

	if o.Side == SELL {
		return profit.Mul(one.Sub(o.FeePercentage))
	}

	return profit.Neg()
}

func (o Order) Cost() big.Decimal {
	cost := o.Amount.Mul(o.Price)

	if o.Side == BUY {
		return cost
	}

	return cost.Neg()
}

// NewOrder returns a new *Order with the designated side
func NewOrder(orderType OrderSide) (o *Order) {
	o = new(Order)
	o.Side = orderType

	return o
}
