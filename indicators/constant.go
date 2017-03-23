package indicators

import (
	. "github.com/sdcoffey/talib4g"
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
