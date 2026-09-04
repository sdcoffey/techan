package techan

// NewCrossUpIndicatorRule returns a new rule that is satisfied when the lower indicator has crossed above the upper
// indicator.
func NewCrossUpIndicatorRule(upper, lower Indicator) Rule {
	return crossRule{
		upper: upper,
		lower: lower,
		cmp:   1,
	}
}

// NewCrossDownIndicatorRule returns a new rule that is satisfied when the upper indicator has crossed below the lower
// indicator.
func NewCrossDownIndicatorRule(upper, lower Indicator) Rule {
	return crossRule{
		upper: lower,
		lower: upper,
		cmp:   -1,
	}
}

type crossRule struct {
	upper Indicator
	lower Indicator
	cmp   int
}

func (cr crossRule) IsSatisfied(index int, record *TradingRecord) bool {
	first := max(FirstValidIndex(cr.upper), FirstValidIndex(cr.lower))
	if index <= first {
		return false
	}
	lower, upper := cr.lower.Calculate(index), cr.upper.Calculate(index)
	if lower.NaN() || upper.NaN() || lower.Cmp(upper) != cr.cmp {
		return false
	}
	for i := index - 1; i >= first; i-- {
		lower, upper = cr.lower.Calculate(i), cr.upper.Calculate(i)
		if lower.NaN() || upper.NaN() {
			return false
		}
		if cmp := lower.Cmp(upper); cmp != 0 {
			return cmp == -cr.cmp
		}
	}
	return false
}
