package talib4g

type relativeStrengthIndicator struct {
	avgGain Indicator
	avgLoss Indicator
}

func NewRelativeStrengthIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndicator{
		avgGain: NewAverageGainsIndicator(indicator, timeframe),
		avgLoss: NewAverageLossesIndicator(indicator, timeframe),
	}
}

func (rsi relativeStrengthIndicator) Calculate(index int) float64 {
	if index == 0 {
		return 0.0
	}

	averageGain := rsi.avgGain.Calculate(index)
	averageLoss := rsi.avgLoss.Calculate(index)

	relativeStrength := 0.0
	if averageLoss > 0 {
		relativeStrength = averageGain / averageLoss
	}

	return 100.0 - (100.0 / (1 + relativeStrength))
}
