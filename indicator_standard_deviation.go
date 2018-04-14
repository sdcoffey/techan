package techan

import (
	"github.com/sdcoffey/big"
)

// StandardDeviationIndicator calculates the standard deviation of a base indicator.
// See https://www.investopedia.com/terms/s/standarddeviation.asp
type StandardDeviationIndicator struct {
	Indicator Indicator
}

// Calculate returns the standard deviation of a base indicator
func (sdi StandardDeviationIndicator) Calculate(index int) big.Decimal {
	return VarianceIndicator{
		Indicator: sdi.Indicator,
	}.Calculate(index).Sqrt()
}
