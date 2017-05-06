package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCandle_SetsBeginTime(t *testing.T) {
	now := time.Now()
	candle := NewCandle(time.Minute, now.Add(time.Minute))
	assert.EqualValues(t, now.Add(time.Minute).UnixNano(), candle.EndTime.UnixNano())
	assert.EqualValues(t, now.UnixNano(), candle.BeginTime.UnixNano())
	assert.EqualValues(t, time.Minute, candle.Period)
}

func TestCandle_AddTrade(t *testing.T) {
	candle := NewCandle(time.Minute, time.Now().Add(time.Minute))

	candle.AddTrade(1, 2) // Open
	candle.AddTrade(1, 5) // High
	candle.AddTrade(1, 1) // Low
	candle.AddTrade(1, 3) // No Diff
	candle.AddTrade(1, 3) // Close

	assert.EqualValues(t, 2, candle.OpenPrice)
	assert.EqualValues(t, 5, candle.MaxPrice)
	assert.EqualValues(t, 1, candle.MinPrice)
	assert.EqualValues(t, 3, candle.ClosePrice)
	assert.EqualValues(t, 5, candle.Volume)
	assert.EqualValues(t, 5, candle.TradeCount)
}
