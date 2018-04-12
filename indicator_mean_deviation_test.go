package techan

import (
	"testing"
)

func TestMeanDeviationIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(1, 2, 7, 6, 3, 4, 5, 11, 3, 0, 9)

	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ts), 5)

	decimalEquals(t, 2.4444, meanDeviation.Calculate(2))
	decimalEquals(t, 2.5, meanDeviation.Calculate(3))
	decimalEquals(t, 2.16, meanDeviation.Calculate(4))
	decimalEquals(t, 1.68, meanDeviation.Calculate(5))
	decimalEquals(t, 1.2, meanDeviation.Calculate(6))
	decimalEquals(t, 2.16, meanDeviation.Calculate(7))
	decimalEquals(t, 2.32, meanDeviation.Calculate(8))
	decimalEquals(t, 2.72, meanDeviation.Calculate(9))
	decimalEquals(t, 3.52, meanDeviation.Calculate(10))
}
