package talib4g

import (
	"github.com/shopspring/decimal"
)

type crossIndicatorRule struct {
	cross CrossIndicator
}

func (cir crossIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return cir.cross.Calculate(index).Cmp(decimal.Zero) > 0
}

func NewCrossUpIndicatorRule(i1, i2 Indicator) Rule {
	return crossIndicatorRule{
		cross: CrossIndicator{i2, i1},
	}
}

func NewCrossDownIndicatorRule(i1, i2 Indicator) Rule {
	return crossIndicatorRule{
		cross: CrossIndicator{i1, i2},
	}
}
