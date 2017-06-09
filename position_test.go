package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPosition_NoOrders_IsNew(t *testing.T) {
	position := newPosition()

	assert.True(t, position.IsNew())
}

func TestPosition_NewPosition_IsOpen(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = NS(1)
	order.Price = NM(2, USD)

	position := NewPosition(order)
	assert.True(t, position.IsOpen())
	assert.False(t, position.IsNew())
	assert.False(t, position.IsClosed())
}

func TestNewPosition_WithBuy_IsLong(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = NS(1)
	order.Price = NM(2, USD)

	position := NewPosition(order)
	assert.True(t, position.IsLong())
}

func TestNewPosition_WithSell_IsShort(t *testing.T) {
	order := NewOrder(SELL)
	order.Amount = NS(1)
	order.Price = NM(2, USD)

	position := NewPosition(order)
	assert.True(t, position.IsShort())
}

func TestPosition_Enter(t *testing.T) {
	position := newPosition()

	order := NewOrder(BUY)
	order.Amount = NS(1)
	order.Price = NM(3, USD)
	order.ExecutionTime = time.Now()

	position.Enter(order)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, order.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, order.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, order.ExecutionTime, position.EntranceOrder().ExecutionTime)
}

func TestPosition_Close(t *testing.T) {
	position := newPosition()

	entranceOrder := NewOrder(BUY)
	entranceOrder.Amount = NS(1)
	entranceOrder.Price = NM(1, USD)
	entranceOrder.ExecutionTime = time.Now()

	position.Enter(entranceOrder)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, entranceOrder.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, entranceOrder.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, entranceOrder.ExecutionTime, position.EntranceOrder().ExecutionTime)

	exitOrder := NewOrder(SELL)
	entranceOrder.Amount = NS(1)
	entranceOrder.Price = NM(4, USD)
	exitOrder.ExecutionTime = time.Now()

	position.Exit(exitOrder)

	assert.True(t, position.IsClosed())

	assert.EqualValues(t, exitOrder.Amount, position.ExitOrder().Amount)
	assert.EqualValues(t, exitOrder.Price, position.ExitOrder().Price)
	assert.EqualValues(t, exitOrder.ExecutionTime, position.ExitOrder().ExecutionTime)
}

func TestPosition_CostBasis(t *testing.T) {
	p := newPosition()

	order := NewOrder(BUY)
	order.Amount = NS(1)
	order.Price = NM(10, USD)

	p.Enter(order)

	costBasis := NM(10, USD)

	assert.EqualValues(t, costBasis.Value(), p.CostBasis().Value())
}

func TestPosition_ExitValue(t *testing.T) {
	p := newPosition()

	order := NewOrder(BUY)
	order.Amount = NS(1)
	order.Price = NM(10, USD)

	p.Enter(order)

	order = NewOrder(SELL)
	order.Amount = NS(1)
	order.Price = NM(12, USD)

	p.Exit(order)

	sellValue := NM(12, USD)

	assert.EqualValues(t, sellValue.Value(), p.ExitValue().Value())
}
