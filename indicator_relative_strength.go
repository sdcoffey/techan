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

func (rsi relativeStrengthIndicator) Calculate(index int) Decimal {
	if index == 0 {
		return ZERO
	}

	averageGain := rsi.avgGain.Calculate(index)
	averageLoss := rsi.avgLoss.Calculate(index)

	relativeStrength := ZERO
	if averageLoss.GT(ZERO) {
		relativeStrength = averageGain.Div(averageLoss)
	}

	oneHundred := TEN.Mul(TEN)
	return oneHundred.Sub(oneHundred.Div(ONE.Add(relativeStrength)))
}
