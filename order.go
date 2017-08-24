package talib4g

import "time"

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)

type order struct {
	Type          OrderSide
	Price         Decimal
	Amount        Decimal
	ExecutionTime time.Time
}

func NewOrder(orderType OrderSide) (o *order) {
	o = new(order)
	o.Type = orderType

	return o
}
