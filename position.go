package talib4g

// A pair of two Order objects
type Position struct {
	orders [2]*order
}

func newPosition() (t *Position) {
	return &Position{
		orders: [2]*order{nil, nil},
	}
}

func NewPosition(openOrder *order) (t *Position) {
	t = new(Position)
	t.orders[0] = openOrder

	return t
}

func (p *Position) Enter(order *order) {
	if order != nil {
		p.orders[0] = order
	}
}

func (p *Position) Exit(order *order) {
	if order != nil {
		p.orders[1] = order
	}
}

func (p *Position) IsLong() bool {
	return p.EntranceOrder() != nil && p.EntranceOrder().Type == BUY
}

func (p *Position) IsShort() bool {
	return p.EntranceOrder() != nil && p.EntranceOrder().Type == SELL
}

func (p *Position) IsOpen() bool {
	return p.EntranceOrder() != nil && p.ExitOrder() == nil
}

func (p *Position) IsClosed() bool {
	return p.EntranceOrder() != nil && p.ExitOrder() != nil
}

func (p *Position) IsNew() bool {
	return p.EntranceOrder() == nil && p.ExitOrder() == nil
}

func (p *Position) EntranceOrder() *order {
	return p.orders[0]
}

func (p *Position) ExitOrder() *order {
	return p.orders[1]
}

func (p *Position) CostBasis() Money {
	if p.IsOpen() || p.IsClosed() {
		return p.EntranceOrder().Amount.Convert(p.EntranceOrder().Price)
	}
	return NM(0, USD)
}

func (p *Position) ExitValue() Money {
	if p.IsClosed() {
		return p.ExitOrder().Amount.Convert(p.ExitOrder().Price)
	}
	return NM(0, USD)
}
