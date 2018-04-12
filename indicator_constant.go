package techan

import "github.com/sdcoffey/big"

type constantIndicator float64

// NewConstantIndicator returns an indicator which always returns the same value for any index. It's useful when combined
// with other, fluxuating indicators to determine when an indicator has crossed a threshold.
func NewConstantIndicator(constant float64) Indicator {
	return constantIndicator(constant)
}

func (ci constantIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(float64(ci))
}
