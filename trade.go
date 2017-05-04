package talib4g

// A pair of two Order objects
type Position struct {
	orders [2]*Order
}

func newTrade() (t *Position) {
	return &Position{
		orders: [2]*Order{nil, nil},
	}
}

func NewTrade(openOrder *Order) (t *Position) {
	t = new(Position)
	t.orders[0] = openOrder

	return t
}

func (this *Position) Enter(order *Order) {
	if order != nil {
		this.orders[0] = order
	}
}

func (this *Position) Exit(order *Order) {
	if order != nil {
		this.orders[1] = order
	}
}

func (this *Position) IsLong() bool {
	return this.EntranceOrder() != nil && this.EntranceOrder().Type == BUY
}

func (this *Position) IsShort() bool {
	return this.EntranceOrder() != nil && this.EntranceOrder().Type == SELL
}

func (this *Position) IsOpen() bool {
	return this.EntranceOrder() != nil && this.ExitOrder() == nil
}

func (this *Position) IsClosed() bool {
	return this.EntranceOrder() != nil && this.ExitOrder() != nil
}

func (this *Position) IsNew() bool {
	return this.EntranceOrder() == nil && this.ExitOrder() == nil
}

func (this *Position) EntranceOrder() *Order {
	return this.orders[0]
}

func (this *Position) ExitOrder() *Order {
	return this.orders[1]
}
