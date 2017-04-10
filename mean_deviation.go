package talib4g

import "github.com/shopspring/decimal"

type MeanDeviationIndicator struct {
	ind       Indicator
	sma       SMAIndicator
	timeFrame int
}

func NewMeanDeviationIndicator(ind Indicator, timeFrame int) MeanDeviationIndicator {
	return MeanDeviationIndicator{
		ind: ind,
		sma: SMAIndicator{
			Indicator: ind,
			TimeFrame: timeFrame,
		},
		timeFrame: timeFrame,
	}
}

func (this MeanDeviationIndicator) Calculate(index int) decimal.Decimal {
	absoluteDeviations := ZERO

	average := this.sma.Calculate(index)
	startIndex := Max(0, index-this.timeFrame+1)

	for i := startIndex; i <= index; i++ {
		absoluteDeviations = absoluteDeviations.Add(this.ind.Calculate(i).Sub(average).Abs())
	}

	return absoluteDeviations.Div(NewDecimal(index - startIndex + 1))
}
