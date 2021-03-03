package techan

import "testing"

func TestExponentialMovingAverage(t *testing.T) {
	expectedValues := []float64{
		0,
		0,
		64.09,
		63.91,
		63.73,
		63.46,
		63.685,
		63.7675,
		63.3588,
		63.3644,
		62.3472,
		61.9286,
	}

	closePriceIndicator := NewClosePriceIndicator(mockedTimeSeries)
	indicatorEquals(t, expectedValues, NewEMAIndicator(closePriceIndicator, 3))
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
