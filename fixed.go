package talib4g

import (
	"github.com/shopspring/decimal"
)

type FixedIndicator []decimal.Decimal

func NewFixedIndicator(vals ...int) FixedIndicator {
	slc := make([]decimal.Decimal, len(vals))
	for i, val := range vals {
		slc[i] = decimal.NewFromFloat(float64(val))
	}

	return FixedIndicator(slc)
}

func (this FixedIndicator) Calculate(index int) decimal.Decimal {
	return this[index]
}
