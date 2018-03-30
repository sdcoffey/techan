package talib4g

import (
	"github.com/sdcoffey/big"
)

type relativeStrengthIndexIndicator struct {
	rsIndicator Indicator
}

func NewRelativeStrengthIndexIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndexIndicator{
		rsIndicator: NewRelativeStrengthIndicator(indicator, timeframe),
	}
}

func (rsi relativeStrengthIndexIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return big.ZERO
	}

	relativeStrength := rsi.rsIndicator.Calculate(index)
	oneHundred := big.NewFromString("100")

	return oneHundred.Sub(oneHundred.Div(big.ONE.Add(relativeStrength)))
}

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

func (rs relativeStrengthIndicator) Calculate(index int) big.Decimal {
	avgGain := rs.avgGain.Calculate(index)
	avgLoss := rs.avgLoss.Calculate(index)

	return avgGain.Div(avgLoss)
}
