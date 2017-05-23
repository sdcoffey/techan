package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeSeries_AddCandle(t *testing.T) {
	t.Run("Throws if nil candle passed", func(t *testing.T) {
		ts := NewTimeSeries()
		assert.Panics(t, func() {
			ts.AddCandle(nil)
		})
	})

	t.Run("Adds candle if last is nil", func(t *testing.T) {
		ts := NewTimeSeries()

		candle := NewCandle(time.Minute, time.Now())
		candle.ClosePrice = 1

		ts.AddCandle(candle)

		assert.Len(t, ts.Candles, 1)
	})

	t.Run("Does not add candle if before last candle", func(t *testing.T) {
		ts := NewTimeSeries()

		now := time.Now()
		candle := NewCandle(time.Minute, now)
		candle.ClosePrice = 1

		ts.AddCandle(candle)

		then := now.Add(-time.Minute * 10)

		nextCandle := NewCandle(time.Minute, then)
		nextCandle.ClosePrice = 2

		ts.AddCandle(nextCandle)

		assert.Len(t, ts.Candles, 1)
		assert.EqualValues(t, now.UnixNano(), ts.Candles[0].EndTime.UnixNano())
	})
}

func TestTimeSeries_LastCandle(t *testing.T) {
	ts := NewTimeSeries()

	now := time.Now()
	candle := NewCandle(time.Minute, now)
	candle.ClosePrice = 1

	ts.AddCandle(candle)

	assert.EqualValues(t, now.UnixNano(), ts.LastCandle().EndTime.UnixNano())
	assert.EqualValues(t, 1, ts.LastCandle().ClosePrice)

	next := time.Now().Add(time.Minute)
	newCandle := NewCandle(time.Minute, next)
	newCandle.ClosePrice = 2

	ts.AddCandle(newCandle)

	assert.EqualValues(t, next.UnixNano(), ts.LastCandle().EndTime.UnixNano())
	assert.EqualValues(t, 2, ts.LastCandle().ClosePrice)
}

func TestTimeSeries_LastIndex(t *testing.T) {
	ts := NewTimeSeries()

	candle := NewCandle(time.Minute, time.Now())
	ts.AddCandle(candle)

	assert.EqualValues(t, 0, ts.LastIndex())

	candle = NewCandle(time.Minute, time.Now())
	ts.AddCandle(candle)

	assert.EqualValues(t, 1, ts.LastIndex())
}
