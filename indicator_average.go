package talib4g

type averageIndicator struct {
	Indicator
	window int
}

// Returns a new average gains indicator, which returns the average gains
// up until that index.
// @param price indicator should not be > 1 derivation removed from a
// timeseries, i.e., a ClosePriceIndicator, VolumeIndicator, etc
func NewAverageGainsIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeGainsIndicator(indicator, window),
		window,
	}
}

// Returns a new average losses indicator, which returns the average losses
// up until that index.
// @param price indicator should not be > 1 derivation removed from a
// timeseries, i.e., a ClosePriceIndicator, VolumeIndicator, etc
func NewAverageLossesIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		NewCumulativeLossesIndicator(indicator, window),
		window,
	}
}

func (ai averageIndicator) Calculate(index int) Decimal {
	return ai.Indicator.Calculate(index).Div(NewDecimal(float64(Min(index+1, ai.window))))
}
