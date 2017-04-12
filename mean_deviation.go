package talib4g

import (
	"math"
)

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

func (this MeanDeviationIndicator) Calculate(index int) float64 {
	absoluteDeviations := 0.0

	average := this.sma.Calculate(index)
	startIndex := Max(0, index-this.timeFrame+1)

	for i := startIndex; i <= index; i++ {
		average += this.ind.Calculate(i) - math.Abs(average)
	}

	return absoluteDeviations / float64(index-startIndex+1)
}
