package techan

import (
	"testing"
)

func TestMeanDeviationIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(1, 2, 7, 6, 3, 4, 5, 11, 3, 0, 9)

	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ts), 5)

	expected := []float64{
		0,
		0,
		0,
		0,
		2.16,
		1.68,
		1.2,
		2.16,
		2.32,
		2.72,
		3.52,
	}

	indicatorEquals(t, expected, meanDeviation)
}
