package techan

import "github.com/sdcoffey/big"

type trueRangeIndicator struct {
	series *TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series *TimeSeries) Indicator {
	return trueRangeIndicator{
		series: series,
	}
}

func (tri trueRangeIndicator) Calculate(index int) big.Decimal {
	if index-1 < 0 {
		return big.ZERO
	}

	candle := tri.series.Candles[index]
	previousClose := tri.series.Candles[index-1].ClosePrice

	trueHigh := big.MaxSlice(candle.MaxPrice, previousClose)
	trueLow := big.MinSlice(candle.MinPrice, previousClose)

	return trueHigh.Sub(trueLow)
}
