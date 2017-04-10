package talib4g

import (
	"github.com/shopspring/decimal"
)

type ConstantIndicator decimal.Decimal

func (this ConstantIndicator) Calculate(index int) decimal.Decimal {
	return decimal.Decimal(this)
}

type OverIndicator struct {
	High Indicator
	Low  Indicator
}
