package talib4g

import "github.com/sdcoffey/big"

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
func (ci crossIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return big.ZERO
	}

	if ci.differenceIndicator.Calculate(index).LTE(big.ZERO) && ci.differenceIndicator.Calculate(index-1).GT(big.ZERO) {
		return big.ONE
	}

	return big.ZERO
}
