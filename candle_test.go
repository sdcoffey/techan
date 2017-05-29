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

	candle.AddTrade(NM(1, security), NM(2, USD)) // Open
	candle.AddTrade(NM(1, security), NM(5, USD)) // High
	candle.AddTrade(NM(1, security), NM(1, USD)) // Low
	candle.AddTrade(NM(1, security), NM(3, USD)) // No Diff
	candle.AddTrade(NM(1, security), NM(3, USD)) // Close

	assert.EqualValues(t, 2, candle.OpenPrice.Float())
	assert.EqualValues(t, 5, candle.MaxPrice.Float())
	assert.EqualValues(t, 1, candle.MinPrice.Float())
	assert.EqualValues(t, 3, candle.ClosePrice.Float())
	assert.EqualValues(t, 5, candle.Volume.Float())
	assert.EqualValues(t, 5, candle.TradeCount)
}
