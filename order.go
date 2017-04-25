package talib4g

import "time"

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)

type Order struct {
	Type          OrderSide
	Price         float64
	Amount        float64
	ExecutionTime time.Time
}

func NewOrder(orderType OrderSide) (o *Order) {
	o = new(Order)
	o.Type = orderType

	return o
}
