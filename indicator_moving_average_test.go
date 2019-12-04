package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMovingAverage(t *testing.T) {
	ts := mockTimeSeriesFl(1, 2, 3, 4, 3, 4, 5, 4, 3, 3, 4, 3, 2)

	sma := NewSimpleMovingAverage(NewClosePriceIndicator(ts), 3)

	decimalEquals(t, 1, sma.Calculate(0))
	decimalEquals(t, 1.5, sma.Calculate(1))

	decimalEquals(t, 2, sma.Calculate(2))
	decimalEquals(t, 3, sma.Calculate(3))
	decimalEquals(t, 10.0/3.0, sma.Calculate(4))
	decimalEquals(t, 11.0/3.0, sma.Calculate(5))
	decimalEquals(t, 4, sma.Calculate(6))
	decimalEquals(t, 13.0/3.0, sma.Calculate(7))
	decimalEquals(t, 4, sma.Calculate(8))
	decimalEquals(t, 10.0/3.0, sma.Calculate(9))
	decimalEquals(t, 10.0/3.0, sma.Calculate(10))
	decimalEquals(t, 10.0/3.0, sma.Calculate(11))
	decimalEquals(t, 3, sma.Calculate(12))
}

func TestExponentialMovingAverage(t *testing.T) {
	t.Run("if index == 0, returns base indicator", func(t *testing.T) {
		ts := mockTimeSeriesFl(
			64.75, 63.79, 63.73,
			63.73, 63.55, 63.19,
			63.91, 63.85, 62.95,
			63.37, 61.33, 61.51)

		ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)

		decimalEquals(t, 64.75, ema.Calculate(0))
	})

	t.Run("if index > 0, calculates ema", func(t *testing.T) {
		ts := mockTimeSeriesFl(
			64.75, 63.79, 63.73,
			63.73, 63.55, 63.19,
			63.91, 63.85, 62.95,
			63.37, 61.33, 61.51)

		ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)

		decimalEquals(t, 63.6948, ema.Calculate(9))
		decimalEquals(t, 63.2649, ema.Calculate(10))
		decimalEquals(t, 62.9458, ema.Calculate(11))
	})

	t.Run("expands result cache when > 10000 candles added", func(t *testing.T) {
		series := randomTimeSeries(10001)
		ema := NewEMAIndicator(NewClosePriceIndicator(series), 10)

		ema.Calculate(10000)

		assert.EqualValues(t, 10001, len(ema.(*emaIndicator).resultCache))
	})
}

func TestNewMMAIndicator(t *testing.T) {
	series := mockTimeSeriesFl(
		64.75, 63.79, 63.73,
		63.73, 63.55, 63.19,
		63.91, 63.85, 62.95,
		63.37, 61.33, 61.51)

	mma := NewMMAIndicator(NewClosePriceIndicator(series), 10)

	decimalEquals(t, 63.9983, mma.Calculate(9))
	decimalEquals(t, 63.7315, mma.Calculate(10))
	decimalEquals(t, 63.5094, mma.Calculate(11))
}

func TestNewMACDIndicator(t *testing.T) {
	series := randomTimeSeries(100)

	macd := NewMACDIndicator(NewClosePriceIndicator(series), 12, 26)

	assert.NotNil(t, macd)
}

func TestNewMACDHistogramIndicator(t *testing.T) {
	series := randomTimeSeries(100)

	macd := NewMACDIndicator(NewClosePriceIndicator(series), 12, 26)
	macdHistogram := NewMACDHistogramIndicator(macd, 9)

	assert.NotNil(t, macdHistogram)
}

func BenchmarkExponetialMovingAverage(b *testing.B) {
	size := 10000
	ts := randomTimeSeries(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)
		ema.Calculate(size - 1)
	}
}
