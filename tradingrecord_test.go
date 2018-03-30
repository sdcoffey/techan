package talib4g

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestNewTradingRecord(t *testing.T) {
	record := NewTradingRecord()

	assert.Len(t, record.Trades, 0)
	assert.True(t, record.CurrentPosition().IsNew())
}

func TestTradingRecord_CurrentTrade(t *testing.T) {
	record := NewTradingRecord()

	yesterday := time.Now().Add(-time.Hour * 24)

	record.Enter(big.NewDecimal(1), big.NewDecimal(2), big.ZERO, example, yesterday)

	assert.EqualValues(t, 1, record.CurrentPosition().EntranceOrder().Price.Float())
	assert.EqualValues(t, 2, record.CurrentPosition().EntranceOrder().Amount.Float())
	assert.EqualValues(t, yesterday.UnixNano(),
		record.CurrentPosition().EntranceOrder().ExecutionTime.UnixNano())

	now := time.Now()
	record.Exit(big.NewDecimal(3), big.NewDecimal(4), big.ZERO, example, now)
	assert.True(t, record.CurrentPosition().IsNew())

	lastTrade := record.LastTrade()

	assert.EqualValues(t, 3, lastTrade.ExitOrder().Price.Float())
	assert.EqualValues(t, 4, lastTrade.ExitOrder().Amount.Float())
	assert.EqualValues(t, now.UnixNano(),
		lastTrade.ExitOrder().ExecutionTime.UnixNano())
}

func TestTradingRecord_Enter(t *testing.T) {
	t.Run("Does not add trades older than last trade", func(t *testing.T) {
		record := NewTradingRecord()

		now := time.Now()
		record.Enter(big.NewDecimal(1), big.NewDecimal(2), big.ZERO, example, now)
		record.Exit(big.NewDecimal(2), big.NewDecimal(2), big.ZERO, example, now.Add(time.Minute))

		record.Enter(big.NewDecimal(2), big.NewDecimal(2), big.ZERO, example, now.Add(-time.Minute))

		assert.True(t, record.CurrentPosition().IsNew())
		assert.Len(t, record.Trades, 1)
	})
}

func TestTradingRecord_Exit(t *testing.T) {
	t.Run("Does not add trades older than last trade", func(t *testing.T) {
		record := NewTradingRecord()

		now := time.Now()
		record.Enter(big.NewDecimal(1), big.NewDecimal(2), big.ZERO, example, now)
		record.Exit(big.NewDecimal(2), big.NewDecimal(2), big.ZERO, example, now.Add(-time.Minute))

		assert.True(t, record.CurrentPosition().IsOpen())
	})
}
