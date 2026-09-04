package techan

import "github.com/sdcoffey/big"

// NewMaximumDrawdownIndicator returns a derivative Indicator which returns the maximum
// drawdown of the underlying indicator over a window. Maximum drawdown is defined as the
// maximum observed loss from peak of an underlying indicator in a given timeframe.
// Maximum drawdown is given as a percentage of the peak. Use a window value of -1 to include
// all values present in the underlying indicator.
// Only declines after positive peaks are measured; if none exist, it returns zero.
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
	start := mdi.FirstValidIndex()
	if mdi.window > 0 {
		start = max(start, index-mdi.window+1)
	}
	if index < start {
		return big.ZERO
	}
	peak := mdi.indicator.Calculate(start)
	drawdown := big.ZERO
	for i := start; i <= index; i++ {
		value := mdi.indicator.Calculate(i)
		if value.NaN() {
			return big.NaN
		}
		if value.GT(peak) {
			peak = value
		}
		// A percentage drawdown needs a positive reference price.
		if peak.GT(big.ZERO) {
			loss := value.Sub(peak).Div(peak)
			if loss.LT(drawdown) {
				drawdown = loss
			}
		}
	}
	return drawdown
}
