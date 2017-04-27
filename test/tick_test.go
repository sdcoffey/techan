package test

import (
	"github.com/sdcoffey/talib4g"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTick_SetsBeginTime(t *testing.T) {
	now := time.Now()
	tick := talib4g.NewTick(time.Minute, now.Add(time.Minute))
	assert.EqualValues(t, now.Add(time.Minute).UnixNano(), tick.EndTime.UnixNano())
	assert.EqualValues(t, now.UnixNano(), tick.BeginTime.UnixNano())
	assert.EqualValues(t, time.Minute, tick.Period)
}

func TestTick_AddTrade(t *testing.T) {
	tick := talib4g.NewTick(time.Minute, time.Now().Add(time.Minute))

	tick.AddTrade(1, 2) // Open
	tick.AddTrade(1, 5) // High
	tick.AddTrade(1, 1) // Low
	tick.AddTrade(1, 3) // Close

	assert.EqualValues(t, 2, tick.OpenPrice)
	assert.EqualValues(t, 5, tick.MaxPrice)
	assert.EqualValues(t, 1, tick.MinPrice)
	assert.EqualValues(t, 3, tick.ClosePrice)

	assert.EqualValues(t, 4, tick.Amount)
	assert.EqualValues(t, 11, tick.Volume)
}
