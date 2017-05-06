package talib4g

type crossIndicatorRule struct {
	cross Indicator
}

func (cir crossIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return cir.cross.Calculate(index) > 0
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
