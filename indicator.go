package talib4g

import (
	"github.com/shopspring/decimal"
)

type Indicator interface {
	Calculate(int) decimal.Decimal
}
