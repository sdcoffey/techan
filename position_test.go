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
	order.Amount = 1
	order.Price = 2

	position := NewPosition(order)
	assert.True(t, position.IsOpen())
	assert.False(t, position.IsNew())
	assert.False(t, position.IsClosed())
}

func TestNewPosition_WithBuy_IsLong(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = 1
	order.Price = 2

	position := NewPosition(order)
	assert.True(t, position.IsLong())
}

func TestNewPosition_WithSell_IsShort(t *testing.T) {
	order := NewOrder(SELL)
	order.Amount = 1
	order.Price = 2

	position := NewPosition(order)
	assert.True(t, position.IsShort())
}

func TestPosition_Enter(t *testing.T) {
	position := newPosition()

	order := NewOrder(BUY)
	order.Amount = 1
	order.Price = 1
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
	entranceOrder.Amount = 1
	entranceOrder.Price = 1
	entranceOrder.ExecutionTime = time.Now()

	position.Enter(entranceOrder)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, entranceOrder.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, entranceOrder.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, entranceOrder.ExecutionTime, position.EntranceOrder().ExecutionTime)

	exitOrder := NewOrder(SELL)
	exitOrder.Amount = 1
	exitOrder.Price = 4
	exitOrder.ExecutionTime = time.Now()

	position.Exit(exitOrder)

	assert.True(t, position.IsClosed())

	assert.EqualValues(t, exitOrder.Amount, position.ExitOrder().Amount)
	assert.EqualValues(t, exitOrder.Price, position.ExitOrder().Price)
	assert.EqualValues(t, exitOrder.ExecutionTime, position.ExitOrder().ExecutionTime)
}
