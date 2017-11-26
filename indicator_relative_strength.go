package talib4g

import "github.com/sdcoffey/big"

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

func (rsi relativeStrengthIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return big.ZERO
	}

	averageGain := rsi.avgGain.Calculate(index)
	averageLoss := rsi.avgLoss.Calculate(index)

	relativeStrength := big.ZERO
	if averageLoss.GT(big.ZERO) {
		relativeStrength = averageGain.Div(averageLoss)
	}

	oneHundred := big.TEN.Frac(10)
	return oneHundred.Sub(oneHundred.Div(big.ONE.Add(relativeStrength)))
}
