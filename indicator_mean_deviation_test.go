package talib4g

import (
	"testing"
)

func TestMeanDeviationIndicator(t *testing.T) {
	ts := MockTimeSeries(1, 2, 7, 6, 3, 4, 5, 11, 3, 0, 9)

	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ts), 5)

	decimalEquals(t, 2.44, meanDeviation.Calculate(2))
	decimalEquals(t, 2.5, meanDeviation.Calculate(3))
	decimalEquals(t, 2.16, meanDeviation.Calculate(7))
	decimalEquals(t, 2.32, meanDeviation.Calculate(8))
	decimalEquals(t, 2.72, meanDeviation.Calculate(9))
}
