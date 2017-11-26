package talib4g

import "github.com/sdcoffey/big"

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

func (p *Position) CostBasis() big.Decimal {
	if p.EntranceOrder() != nil {
		return p.EntranceOrder().Amount.Mul(p.EntranceOrder().Price)
	}
	return big.NewDecimal(0)
}

func (p *Position) ExitValue() big.Decimal {
	if p.IsClosed() {
		return p.ExitOrder().Amount.Mul(p.ExitOrder().Price)
	}
	return big.NewDecimal(0)
}
