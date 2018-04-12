package techan

import "github.com/sdcoffey/big"

// Position is a pair of two Order objects
type Position struct {
	orders [2]*Order
}

// NewPosition returns a new Position with the passed-in order as the open order
func NewPosition(openOrder Order) (t *Position) {
	t = new(Position)
	t.orders[0] = &openOrder

	return t
}

// Enter sets the open order to the order passed in
func (p *Position) Enter(order Order) {
	p.orders[0] = &order
}

// Exit sets the exit order to the order passed in
func (p *Position) Exit(order Order) {
	p.orders[1] = &order
}

// IsLong returns true if the entrance order is a buy order
func (p *Position) IsLong() bool {
	return p.EntranceOrder() != nil && p.EntranceOrder().Side == BUY
}

// IsShort returns true if the entrance order is a sell order
func (p *Position) IsShort() bool {
	return p.EntranceOrder() != nil && p.EntranceOrder().Side == SELL
}

// IsOpen returns true if there is an entrance order but no exit order
func (p *Position) IsOpen() bool {
	return p.EntranceOrder() != nil && p.ExitOrder() == nil
}

// IsClosed returns true of there are both entrance and exit orders
func (p *Position) IsClosed() bool {
	return p.EntranceOrder() != nil && p.ExitOrder() != nil
}

// IsNew returns true if there is neither an entrance or exit order
func (p *Position) IsNew() bool {
	return p.EntranceOrder() == nil && p.ExitOrder() == nil
}

// EntranceOrder returns the entrance order of this position
func (p *Position) EntranceOrder() *Order {
	return p.orders[0]
}

// ExitOrder returns the exit order of this position
func (p *Position) ExitOrder() *Order {
	return p.orders[1]
}

// CostBasis returns the price to enter this order
func (p *Position) CostBasis() big.Decimal {
	if p.EntranceOrder() != nil {
		return p.EntranceOrder().Amount.Mul(p.EntranceOrder().Price)
	}
	return big.ZERO
}

// ExitValue returns the value accrued by closing the position
func (p *Position) ExitValue() big.Decimal {
	if p.IsClosed() {
		return p.ExitOrder().Amount.Mul(p.ExitOrder().Price)
	}

	return big.ZERO
}
