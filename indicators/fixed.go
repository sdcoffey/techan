package indicators

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
)

type FixedIndicator []decimal.Decimal

func NewFixedIndicator(vals ...int) FixedIndicator {
	slc := make([]decimal.Decimal, len(vals))
	for i, val := range vals {
		slc[i] = NewDecimal(val)
	}

	return FixedIndicator(slc)
}

func (this FixedIndicator) Calculate(index int) decimal.Decimal {
	return this[index]
}
