package techan

import (
	"testing"
)

func TestTrueRangeIndicator(t *testing.T) {
	trueRangeIndicator := NewTrueRangeIndicator(mockedTimeSeries)

	expectedValues := []float64{
		0,
		2,
		2,
		2,
		2,
		2,
		2,
		2,
		2,
		2,
		3.04,
		2,
	}

	indicatorEquals(t, expectedValues, trueRangeIndicator)
}
