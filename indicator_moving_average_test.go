package techan

import (
	"math"
	"testing"
	"math/big"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMovingAverage(t *testing.T) {
	//ts := mockTimeSeriesFl(15, 16, 17, 18, 17, 18, 16)
	//
	//sma := NewSimpleMovingAverage(NewClosePriceIndicator(ts), 3)

	nan := math.NaN()
	println(nan)

	bnan := big.NewFloat(nan)
	print(bnan.String())
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

	t.Run("Very simple", func(t *testing.T) {
		ts := mockTimeSeriesFl(15, 16, 17, 18, 19, 20, 21)

		ema := NewEMAIndicator(NewClosePriceIndicator(ts), 3)

		decimalEquals(t, 16, ema.Calculate(5))
		decimalEquals(t, 15, ema.Calculate(4))
		decimalEquals(t, 14, ema.Calculate(3))
		decimalEquals(t, 13, ema.Calculate(2))
		decimalEquals(t, 0, ema.Calculate(1))
		decimalEquals(t, 0, ema.Calculate(0))
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
