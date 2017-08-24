package talib4g

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCandle_AddTrade(t *testing.T) {
	now := time.Now()
	candle := NewCandle(NewTimePeriod(now, now.Add(time.Minute)))

	candle.AddTrade(NewDecimal(1), NewDecimal(2)) // Open
	candle.AddTrade(NewDecimal(1), NewDecimal(5)) // High
	candle.AddTrade(NewDecimal(1), NewDecimal(1)) // Low
	candle.AddTrade(NewDecimal(1), NewDecimal(3)) // No Diff
	candle.AddTrade(NewDecimal(1), NewDecimal(3)) // Close

	assert.EqualValues(t, 2, candle.OpenPrice.Float())
	assert.EqualValues(t, 5, candle.MaxPrice.Float())
	assert.EqualValues(t, 1, candle.MinPrice.Float())
	assert.EqualValues(t, 3, candle.ClosePrice.Float())
	assert.EqualValues(t, 5, candle.Volume.Float())
	assert.EqualValues(t, 5, candle.TradeCount)
}
