package techan

import "github.com/sdcoffey/big"

type averageIndicator struct {
	Indicator
	window int
}

// NewAverageGainsIndicator Returns a new average gains indicator, which returns the average gains
// in the given window based on the given indicator.
func NewAverageGainsIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeGainsIndicator(indicator, window),
		window,
	}
}

// NewAverageLossesIndicator Returns a new average losses indicator, which returns the average losses
// in the given window based on the given indicator.
func NewAverageLossesIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeLossesIndicator(indicator, window),
		window,
	}
}

func (ai averageIndicator) Calculate(index int) big.Decimal {
	return ai.Indicator.Calculate(index).Div(big.NewDecimal(float64(Min(index+1, ai.window))))
}
