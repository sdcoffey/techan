package talib4g

import "github.com/sdcoffey/big"

type crossIndicatorRule struct {
	cross Indicator
}

func (cir crossIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return cir.cross.Calculate(index).GT(big.ZERO)
}

func NewCrossUpIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: NewCrossIndicator(upper, lower),
	}
}

func NewCrossDownIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: NewCrossIndicator(lower, upper),
	}
}
