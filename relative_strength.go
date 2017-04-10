package talib4g

import (
	"github.com/shopspring/decimal"
)

type RSIIndicator struct {
	AvgGainIndicator AverageIndicator
	AvgLossIndicator AverageIndicator
	TimeFrame        int
}

func NewRSIIndicator(ind Indicator, timeFrame int) RSIIndicator {
	return RSIIndicator{
		AvgGainIndicator: AverageIndicator{CumulativeGainsIndicator{ind, timeFrame}, timeFrame},
		AvgLossIndicator: AverageIndicator{CumulativeLossesIndicator{ind, timeFrame}, timeFrame},
		TimeFrame:        timeFrame,
	}
}

func (this RSIIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return ZERO
	}

	averageGain := this.AvgGainIndicator.Calculate(index)
	averageLoss := this.AvgLossIndicator.Calculate(index)

	relativeStrength := decimal.Zero
	if averageLoss.Cmp(decimal.Zero) != 0 {
		relativeStrength = averageGain.Div(averageLoss)
	}

	hundred := TEN.Mul(TEN)
	return hundred.Sub(hundred.Div(ONE.Add(relativeStrength)))
}
