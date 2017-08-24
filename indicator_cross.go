package talib4g

type crossIndicator struct {
	differenceIndicator
}

// Returns a new CrossIndicator, which, given an index, determines whether a lower
// indicator has crossed an upper one
func NewCrossIndicator(upper, lower Indicator) Indicator {
	return crossIndicator{
		differenceIndicator{
			minuend:    upper,
			subtrahend: lower,
		},
	}
}

// Walks backward from the current index to determine if the lower indicator
// has crossed the upper indicator. Returns truthy value if so, zero otherwise
func (ci crossIndicator) Calculate(index int) Decimal {
	if index == 0 {
		return ZERO
	}

	if ci.differenceIndicator.Calculate(index).LTE(ZERO) && ci.differenceIndicator.Calculate(index-1).GT(ZERO) {
		return ONE
	}

	return ZERO
}
