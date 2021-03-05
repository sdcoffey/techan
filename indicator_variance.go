package techan

import "github.com/sdcoffey/big"

// NewVarianceIndicator provides a way to find the variance in a base indicator, where variances is the sum of squared
// deviations from the mean at any given index in the time series.
func NewVarianceIndicator(ind Indicator) Indicator {
	return varianceIndicator{
		indicator: ind,
	}
}

type varianceIndicator struct {
	indicator Indicator
}

// Calculate returns the Variance for this indicator at the given index
func (vi varianceIndicator) Calculate(index int) big.Decimal {
	if index < 1 {
		return big.ZERO
	}

	avgIndicator := NewSimpleMovingAverage(vi.indicator, index+1)
	avg := avgIndicator.Calculate(index)
	variance := big.ZERO

	for i := 0; i <= index; i++ {
		pow := vi.indicator.Calculate(i).Sub(avg).Pow(2)
		variance = variance.Add(pow)
	}

	return variance.Div(big.NewDecimal(float64(index + 1)))
}

func (vi varianceIndicator) RemoveCachedEntry(index int) {
	vi.indicator.RemoveCachedEntry(index)
}
