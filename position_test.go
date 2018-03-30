package talib4g

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestPosition_NoOrders_IsNew(t *testing.T) {
	position := new(Position)

	assert.True(t, position.IsNew())
}

func TestPosition_NewPosition_IsOpen(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsOpen())
	assert.False(t, position.IsNew())
	assert.False(t, position.IsClosed())
}

func TestNewPosition_WithBuy_IsLong(t *testing.T) {
	order := NewOrder(BUY)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsLong())
}

func TestNewPosition_WithSell_IsShort(t *testing.T) {
	order := NewOrder(SELL)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(2)

	position := NewPosition(order)
	assert.True(t, position.IsShort())
}

func TestPosition_Enter(t *testing.T) {
	position := new(Position)

	order := NewOrder(BUY)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(3)
	order.ExecutionTime = time.Now()

	position.Enter(order)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, order.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, order.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, order.ExecutionTime, position.EntranceOrder().ExecutionTime)
}

func TestPosition_Close(t *testing.T) {
	position := new(Position)

	entranceOrder := NewOrder(BUY)
	entranceOrder.Amount = big.NewDecimal(1)
	entranceOrder.Price = big.NewDecimal(1)
	entranceOrder.ExecutionTime = time.Now()

	position.Enter(entranceOrder)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, entranceOrder.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, entranceOrder.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, entranceOrder.ExecutionTime, position.EntranceOrder().ExecutionTime)

	exitOrder := NewOrder(SELL)
	entranceOrder.Amount = big.NewDecimal(1)
	entranceOrder.Price = big.NewDecimal(4)
	exitOrder.ExecutionTime = time.Now()

	position.Exit(exitOrder)

	assert.True(t, position.IsClosed())

	assert.EqualValues(t, exitOrder.Amount, position.ExitOrder().Amount)
	assert.EqualValues(t, exitOrder.Price, position.ExitOrder().Price)
	assert.EqualValues(t, exitOrder.ExecutionTime, position.ExitOrder().ExecutionTime)
}

func TestPosition_CostBasis(t *testing.T) {
	p := new(Position)

	order := NewOrder(BUY)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(10)

	p.Enter(order)

	costBasis := big.NewDecimal(10)

	assert.EqualValues(t, costBasis.Float(), p.CostBasis().Float())
}

func TestPosition_ExitValue(t *testing.T) {
	p := new(Position)

	order := NewOrder(BUY)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(10)

	p.Enter(order)

	order = NewOrder(SELL)
	order.Amount = big.NewDecimal(1)
	order.Price = big.NewDecimal(12)

	p.Exit(order)

	sellValue := big.NewDecimal(12)

	assert.EqualValues(t, sellValue.Float(), p.ExitValue().Float())
}
