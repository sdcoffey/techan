package talib4g

import "github.com/sdcoffey/big"

type crossIndicatorRule struct {
	cross Indicator
}

func (cir crossIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return cir.cross.Calculate(index).GT(big.ZERO)
}

// NewCrossUpIndicatorRule returns a new rule that is satisfied when the lower indicator has crossed above the upper
// indicator.
func NewCrossUpIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: NewCrossIndicator(upper, lower),
	}
}

// NewCrossDownIndicatorRule returns a new rule that is satisfied when the upper indicator has crossed below the lower
// indicator.
func NewCrossDownIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: NewCrossIndicator(lower, upper),
	}
}
