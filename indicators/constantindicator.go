package indicators

import (
	"github.com/shopspring/decimal"
)

type ConstantIndicator decimal.Decimal

func (this ConstantIndicator) Calculate(index int) decimal.Decimal {
	return decimal.Decimal(this)
}
