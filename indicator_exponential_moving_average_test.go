package techan

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestExponentialMovingAverage(t *testing.T) {
	t.Run("Default Case", func(t *testing.T) {
		expectedValues := []float64{
			0,
			0,
			0,
			64,
			63.82,
			63.568,
			63.7048,
			63.7629,
			63.4377,
			63.4106,
			62.5784,
			62.151,
		}

		closePriceIndicator := NewClosePriceIndicator(mockedTimeSeries)
		indicatorEquals(t, expectedValues, NewEMAIndicator(closePriceIndicator, 4))
	})

	t.Run("Expands Result Cache", func(t *testing.T) {
		closeIndicator := NewClosePriceIndicator(randomTimeSeries(1001))
		ema := NewEMAIndicator(closeIndicator, 20)

		ema.Calculate(1000)

		emaStruct, ok := ema.(cachedIndicator)
		assert.True(t, ok)
		assert.EqualValues(t, 1001, len(emaStruct.cache()))
	})

	t.Run("Can reset cached values after the underlying series changes", func(t *testing.T) {
		series := mockTimeSeriesFl(10, 10, 10, 10)
		ema := NewEMAIndicator(NewClosePriceIndicator(series), 3)

		decimalEquals(t, 10, ema.Calculate(3))

		series.Candles[3].ClosePrice = big.NewFromString("20")
		decimalEquals(t, 10, ema.Calculate(3))

		assert.True(t, ResetCacheFrom(ema, 3))
		decimalEquals(t, 15, ema.Calculate(3))
	})

	t.Run("Reports when an indicator does not support cache resets", func(t *testing.T) {
		series := mockTimeSeriesFl(10, 20, 30)
		sma := NewSimpleMovingAverage(NewClosePriceIndicator(series), 3)

		assert.False(t, ResetCacheFrom(sma, 0))
	})
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

func BenchmarkExponentialMovingAverage_Cached(b *testing.B) {
	size := 10000
	ts := randomTimeSeries(size)
	ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ema.Calculate(size - 1)
	}
}
