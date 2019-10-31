package techan

import (
	"github.com/sdcoffey/big"
)

// NewStandardDeviationIndicator calculates the standard deviation of a base indicator.
// See https://www.investopedia.com/terms/s/standarddeviation.asp
func NewStandardDeviationIndicator(ind Indicator) Indicator {
	return standardDeviationIndicator{
		indicator: NewVarianceIndicator(ind),
	}
}

type standardDeviationIndicator struct {
	indicator Indicator
}

// Calculate returns the standard deviation of a base indicator
func (sdi standardDeviationIndicator) Calculate(index int) big.Decimal {
	return sdi.indicator.Calculate(index).Sqrt()
}
