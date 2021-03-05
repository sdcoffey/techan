package techan

import "github.com/sdcoffey/big"

type averageTrueRangeIndicator struct {
	series *TimeSeries
	window int
}

// NewAverageTrueRangeIndicator returns a base indicator that calculates the average true range of the
// underlying over a window
// https://www.investopedia.com/terms/a/atr.asp
func NewAverageTrueRangeIndicator(series *TimeSeries, window int) Indicator {
	return averageTrueRangeIndicator{
		series: series,
		window: window,
	}
}

func (atr averageTrueRangeIndicator) Calculate(index int) big.Decimal {
	if index < atr.window {
		return big.ZERO
	}

	sum := big.ZERO

	for i := index; i > index-atr.window; i-- {
		sum = sum.Add(NewTrueRangeIndicator(atr.series).Calculate(i))
	}

	return sum.Div(big.NewFromInt(atr.window))
}
