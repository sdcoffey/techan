package techan

import (
	"testing"
)

func TestModifiedMovingAverage(t *testing.T) {
	indicator := NewMMAIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expected := []float64{
		0,
		0,
		64.09,
		63.97,
		63.83,
		63.6167,
		63.7144,
		63.7596,
		63.4898,
		63.4498,
		62.7432,
		62.3321,
	}

	indicatorEquals(t, expected, indicator)
}
