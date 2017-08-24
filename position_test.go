package talib4g

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPosition_NoOrders_IsNew(t *testing.T) {
	position := newPosition()

	assert.True(t, position.IsNew())
}

func TestPosition_NewPosition_IsOpen(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsOpen())
	assert.False(t, position.IsNew())
	assert.False(t, position.IsClosed())
}

func TestNewPosition_WithBuy_IsLong(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsLong())
}

func TestNewPosition_WithSell_IsShort(t *testing.T) {
	order := NewOrder(SELL)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsShort())
}

func TestPosition_Enter(t *testing.T) {
	position := newPosition()

	order := NewOrder(BUY)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(3)
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
	entranceOrder.Amount = NewDecimal(1)
	entranceOrder.Price = NewDecimal(1)
	entranceOrder.ExecutionTime = time.Now()

	position.Enter(entranceOrder)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, entranceOrder.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, entranceOrder.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, entranceOrder.ExecutionTime, position.EntranceOrder().ExecutionTime)

	exitOrder := NewOrder(SELL)
	entranceOrder.Amount = NewDecimal(1)
	entranceOrder.Price = NewDecimal(4)
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
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(10)

	p.Enter(order)

	costBasis := NewDecimal(10)

	assert.EqualValues(t, costBasis.Float(), p.CostBasis().Float())
}

func TestPosition_ExitValue(t *testing.T) {
	p := newPosition()

	order := NewOrder(BUY)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(10)

	p.Enter(order)

	order = NewOrder(SELL)
	order.Amount = NewDecimal(1)
	order.Price = NewDecimal(12)

	p.Exit(order)

	sellValue := NewDecimal(12)

	assert.EqualValues(t, sellValue.Float(), p.ExitValue().Float())
}
