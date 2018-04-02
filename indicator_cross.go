package talib4g

import "github.com/sdcoffey/big"

type crossIndicator struct {
	upper Indicator
	lower Indicator
}

// NewCrossIndicator returns an Indicator, which, given an index, determines whether a lower
// indicator has crossed an upper one
func NewCrossIndicator(upper, lower Indicator) Indicator {
	return crossIndicator{
		upper: upper,
		lower: lower,
	}
}

// Calculate walks backward from the current index to determine if the lower indicator
// has crossed the upper indicator. Returns truthy value if so, zero otherwise
func (ci crossIndicator) Calculate(index int) big.Decimal {
	i := index

	if i == 0 || ci.upper.Calculate(i).GTE(ci.lower.Calculate(i)) {
		return big.ZERO
	}

	i--

	if ci.upper.Calculate(i).GT(ci.lower.Calculate(i)) {
		return big.ONE
	}

	for i > 0 && ci.upper.Calculate(i).EQ(ci.lower.Calculate(i)) {
		i--
	}

	if i != 0 && ci.upper.Calculate(i).GT(ci.lower.Calculate(i)) {
		return big.ONE
	}

	return big.ZERO
}
