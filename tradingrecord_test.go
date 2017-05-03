package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTradingRecord(t *testing.T) {
	record := NewTradingRecord()

	assert.Len(t, record.Trades, 0)
	assert.True(t, record.CurrentTrade().IsNew())
}

func TestTradingRecord_CurrentTrade(t *testing.T) {
	record := NewTradingRecord()

	now := time.Now()
	yesterday := now.Add(-time.Hour * 24)

	record.Enter(1, 2, yesterday)

	assert.EqualValues(t, 1, record.CurrentTrade().EntranceOrder().Price)
	assert.EqualValues(t, 2, record.CurrentTrade().EntranceOrder().Amount)
	assert.EqualValues(t, yesterday.UnixNano(),
		record.CurrentTrade().EntranceOrder().ExecutionTime.UnixNano())

	record.Exit(3, 4, yesterday.Add(1))
	assert.True(t, record.CurrentTrade().IsNew())

	lastTrade := record.Trades[len(record.Trades)-1]

	assert.EqualValues(t, 3, lastTrade.ExitOrder().Price)
	assert.EqualValues(t, 4, lastTrade.ExitOrder().Amount)
	assert.EqualValues(t, yesterday.Add(1).UnixNano(),
		lastTrade.ExitOrder().ExecutionTime.UnixNano())
}
