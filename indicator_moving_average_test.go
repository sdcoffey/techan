package talib4g

import (
	"testing"
)

func TestSimpleMovingAverage(t *testing.T) {
	ts := MockTimeSeries(1, 2, 3, 4, 3, 4, 5, 4, 3, 3, 4, 3, 2)

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
	ts := MockTimeSeries(
		64.75, 63.79, 63.73,
		63.73, 63.55, 63.19,
		63.91, 63.85, 62.95,
		63.37, 61.33, 61.51)

	ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)

	decimalEquals(t, 63.6536, ema.Calculate(9))
	decimalEquals(t, 63.2312, ema.Calculate(10))
	decimalEquals(t, 62.9182, ema.Calculate(11))
}

func BenchmarkExponetialMovingAverage(b *testing.B) {
	size := 10000
	ts := RandomTimeSeries(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ema := NewEMAIndicator(NewClosePriceIndicator(ts), 10)
		ema.Calculate(size - 1)
	}
}
