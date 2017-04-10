package talib4g

import (
	"github.com/shopspring/decimal"
)

type CrossIndicator struct {
	Upper Indicator
	Lower Indicator
}

func (this CrossIndicator) Calculate(index int) decimal.Decimal {
	i := index

	if i == 0 || this.Upper.Calculate(index).Cmp(this.Lower.Calculate(i)) >= 0 {
		return ZERO
	}

	for ; i > 0; i-- {
		if this.Upper.Calculate(i).Cmp(this.Lower.Calculate(i)) > 0 {
			return ONE
		}
	}

	return ZERO
}
