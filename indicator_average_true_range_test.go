package techan

import "testing"

func TestAverageTrueRangeIndicator(t *testing.T) {
	atrIndicator := NewAverageTrueRangeIndicator(mockedTimeSeries, 3)

	expectedValues := []float64{
		0,
		0,
		0,
		2,
		2,
		2,
		2,
		2,
		2,
		2,
		2.3467,
		2.3467,
	}

	indicatorEquals(t, expectedValues, atrIndicator)
}
