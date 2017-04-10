package talib4g

// A pair of two Order objects
type Trade struct {
	orders [2]*Order
}

func newTrade() (t *Trade) {
	return &Trade{
		orders: [2]*Order{nil, nil},
	}
}

func NewTrade(openOrder *Order) (t *Trade) {
	t = new(Trade)
	t.orders[0] = openOrder

	return t
}

func (this *Trade) Enter(order *Order) {
	if order != nil {
		this.orders[0] = order
	}
}

func (this *Trade) Exit(order *Order) {
	if order != nil {
		this.orders[1] = order
	}
}

func (this *Trade) IsOpen() bool {
	return this.EntranceOrder() != nil && this.ExitOrder() == nil
}

func (this *Trade) IsClosed() bool {
	return this.EntranceOrder() != nil && this.ExitOrder() != nil
}

func (this *Trade) IsNew() bool {
	return this.EntranceOrder() == nil && this.ExitOrder() == nil
}

func (this *Trade) EntranceOrder() *Order {
	return this.orders[0]
}

func (this *Trade) ExitOrder() *Order {
	return this.orders[1]
}
