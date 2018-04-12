package techan

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
	order := Order{
		Side:   BUY,
		Amount: big.ONE,
		Price:  big.NewFromString("2"),
	}

	position := NewPosition(order)
	assert.True(t, position.IsOpen())
	assert.False(t, position.IsNew())
	assert.False(t, position.IsClosed())
}

func TestNewPosition_WithBuy_IsLong(t *testing.T) {
	order := Order{
		Side:   BUY,
		Amount: big.ONE,
		Price:  big.NewFromString("2"),
	}

	position := NewPosition(order)
	assert.True(t, position.IsLong())
}

func TestNewPosition_WithSell_IsShort(t *testing.T) {
	order := Order{
		Side:   SELL,
		Amount: big.ONE,
		Price:  big.NewFromString("2"),
	}

	position := NewPosition(order)
	assert.True(t, position.IsShort())
}

func TestPosition_Enter(t *testing.T) {
	position := new(Position)

	order := Order{
		Side:   BUY,
		Amount: big.ONE,
		Price:  big.NewFromString("2"),
	}

	position.Enter(order)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, order.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, order.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, order.ExecutionTime, position.EntranceOrder().ExecutionTime)
}

func TestPosition_Close(t *testing.T) {
	position := new(Position)

	entranceOrder := Order{
		Side:   BUY,
		Amount: big.ONE,
		Price:  big.NewFromString("2"),
	}

	position.Enter(entranceOrder)

	assert.True(t, position.IsOpen())
	assert.EqualValues(t, entranceOrder.Amount, position.EntranceOrder().Amount)
	assert.EqualValues(t, entranceOrder.Price, position.EntranceOrder().Price)
	assert.EqualValues(t, entranceOrder.ExecutionTime, position.EntranceOrder().ExecutionTime)

	exitOrder := Order{
		Side:          SELL,
		Amount:        big.ONE,
		Price:         big.NewFromString("4"),
		ExecutionTime: time.Now(),
	}

	position.Exit(exitOrder)

	assert.True(t, position.IsClosed())

	assert.EqualValues(t, exitOrder.Amount, position.ExitOrder().Amount)
	assert.EqualValues(t, exitOrder.Price, position.ExitOrder().Price)
	assert.EqualValues(t, exitOrder.ExecutionTime, position.ExitOrder().ExecutionTime)
}

func TestPosition_CostBasis(t *testing.T) {
	t.Run("When entrance order nil, returns 0", func(t *testing.T) {
		p := new(Position)
		assert.EqualValues(t, "0", p.CostBasis().String())
	})

	t.Run("When entracne order not nil, returns cost basis", func(t *testing.T) {
		p := new(Position)

		order := Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.NewFromString("2"),
		}

		p.Enter(order)

		assert.EqualValues(t, "2.00", p.CostBasis().FormattedString(2))
	})
}

func TestPosition_ExitValue(t *testing.T) {
	t.Run("when not closed, returns 0", func(t *testing.T) {
		p := new(Position)

		order := Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.NewFromString("2"),
		}

		p.Enter(order)

		assert.EqualValues(t, "0.00", p.ExitValue().FormattedString(2))
	})

	t.Run("when closed, returns exit value", func(t *testing.T) {
		p := new(Position)

		order := Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.NewFromString("2"),
		}

		p.Enter(order)

		order = Order{
			Side:   SELL,
			Amount: big.ONE,
			Price:  big.NewFromString("12"),
		}

		p.Exit(order)

		assert.EqualValues(t, "12.00", p.ExitValue().FormattedString(2))
	})
}
