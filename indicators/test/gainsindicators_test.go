package test

import (
	. "github.com/sdcoffey/talib4g/indicators"
	"testing"
)

func TestCumulativeGainsIndicator(t *testing.T) {
	ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

	cumGains := CumulativeGainsIndicator{
		Indicator: ClosePriceIndicator(*ts),
		TimeFrame: 10,
	}

	decimalEquals(t, 0, cumGains.Calculate(0))
	decimalEquals(t, 1, cumGains.Calculate(1))
	decimalEquals(t, 2, cumGains.Calculate(2))
	decimalEquals(t, 4, cumGains.Calculate(3))
	decimalEquals(t, 7, cumGains.Calculate(4))
	decimalEquals(t, 12, cumGains.Calculate(5))
}

func TestCumulativeLossesIndicator(t *testing.T) {
	ts := MockTimeSeries(13, 8, 5, 3, 2, 1)

	cumGains := CumulativeLossesIndicator{
		Indicator: ClosePriceIndicator(*ts),
		TimeFrame: 10,
	}

	decimalEquals(t, 0, cumGains.Calculate(0))
	decimalEquals(t, -5, cumGains.Calculate(1))
	decimalEquals(t, -8, cumGains.Calculate(2))
	decimalEquals(t, -10, cumGains.Calculate(3))
	decimalEquals(t, -11, cumGains.Calculate(4))
	decimalEquals(t, -12, cumGains.Calculate(5))
}

func TestAverageIndicator(t *testing.T) {
	ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

	avgGains := AverageIndicator{
		Indicator: CumulativeGainsIndicator{
			Indicator: ClosePriceIndicator(*ts),
			TimeFrame: 10,
		},
		TimeFrame: 10,
	}

	decimalEquals(t, 0, avgGains.Calculate(0))
	decimalEquals(t, 1.0/2.0, avgGains.Calculate(1))
	decimalEquals(t, 2.0/3.0, avgGains.Calculate(2))
	decimalEquals(t, 4.0/4.0, avgGains.Calculate(3))
	decimalEquals(t, 7.0/5.0, avgGains.Calculate(4))
	decimalEquals(t, 12.0/6.0, avgGains.Calculate(5))
}
