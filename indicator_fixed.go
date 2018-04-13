package techan

import "github.com/sdcoffey/big"

type fixedIndicator []float64

// NewFixedIndicator returns an indicator with a fixed set of values that are returned when an index is passed in
func NewFixedIndicator(vals ...float64) Indicator {
	return fixedIndicator(vals)
}

func (fi fixedIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(fi[index])
}
