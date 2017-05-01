package talib4g

type crossIndicatorRule struct {
	cross CrossIndicator
}

func (cir crossIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return cir.cross.Calculate(index) > 0
}

func NewCrossUpIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: CrossIndicator{upper, lower},
	}
}

func NewCrossDownIndicatorRule(upper, lower Indicator) Rule {
	return crossIndicatorRule{
		cross: CrossIndicator{lower, upper},
	}
}
