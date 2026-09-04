package techan

import "github.com/sdcoffey/big"

type averageIndicator struct {
	Indicator
	window int
}

// NewAverageGainsIndicator Returns a new average gains indicator, which returns the average gains
// in the given window based on the given indicator. The divisor is the number
// of usable source observations so far, capped at window.
func NewAverageGainsIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeGainsIndicator(indicator, window),
		window,
	}
}

// NewAverageLossesIndicator Returns a new average losses indicator, which returns the average losses
// in the given window based on the given indicator. The divisor is the number
// of usable source observations so far, capped at window.
func NewAverageLossesIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeLossesIndicator(indicator, window),
		window,
	}
}

func (ai averageIndicator) Calculate(index int) big.Decimal {
	first := ai.FirstValidIndex()
	if index < first {
		return big.ZERO
	}
	// The first cumulative change requires two usable source observations.
	observations := min(index-first+2, ai.window)
	return ai.Indicator.Calculate(index).Div(big.NewFromInt(observations))
}
