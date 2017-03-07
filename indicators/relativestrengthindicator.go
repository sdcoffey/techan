package indicators

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
)

type RSIIndicator struct {
	AvgGainIndicator AverageIndicator
	AvgLossIndicator AverageIndicator
	TimeFrame        int
}

func NewRSIIndicator(ind Indicator, timeFrame int) RSIIndicator {
	return RSIIndicator{
		AvgGainIndicator: AverageIndicator{CumulativeGainsIndicator{ind, timeFrame}},
		AvgLossIndicator: AverageIndicator{CumulativeLossesIndicator{ind, timeFrame}},
		TimeFrame:        timeFrame,
	}
}

func (this RSIIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return ZERO
	}

	averageGain := this.AvgGainIndicator.Calculate(index)
	averageLoss := this.AvgLossIndicator.Calculate(index)
	relativeStrength := averageGain.Div(averageLoss)

	hundred := TEN.Mul(TEN)
	return hundred.Sub(hundred.Div(ONE.Add(relativeStrength)))
}
