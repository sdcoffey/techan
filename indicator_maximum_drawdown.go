package techan

import "github.com/sdcoffey/big"

// NewMaximumDrawdownIndicator returns a derivative Indicator which returns the maximum
// drawdown of the underlying indicator over a window. Maximum drawdown is defined as the
// maximum observed loss from peak of an underlying indicator in a given timeframe.
// Maximum drawdown is given as a percentage of the peak. Use a window value of -1 to include
// all values present in the underlying indicator.
// See: https://www.investopedia.com/terms/m/maximum-drawdown-mdd.asp
func NewMaximumDrawdownIndicator(ind Indicator, window int) Indicator {
	return maximumDrawdownIndicator{
		indicator: ind,
		window:    window,
	}
}

type maximumDrawdownIndicator struct {
	indicator Indicator
	window    int
}

func (mdi maximumDrawdownIndicator) Calculate(index int) big.Decimal {
	minVal := NewMinimumValueIndicator(mdi.indicator, mdi.window).Calculate(index)
	maxVal := NewMaximumValueIndicator(mdi.indicator, mdi.window).Calculate(index)

	return (minVal.Sub(maxVal)).Div(maxVal)
}
