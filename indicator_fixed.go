package talib4g

import "github.com/sdcoffey/big"

type fixedIndicator []float64

// NewFixedIndicator returns an indicator with a fixed set of values that are returned when an index is passed in
func NewFixedIndicator(vals ...float64) Indicator {
	slc := make([]float64, len(vals))
	for i, val := range vals {
		slc[i] = val
	}

	return fixedIndicator(slc)
}

func (fi fixedIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(fi[index])
}
