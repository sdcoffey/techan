package talib4g

import (
	"math"
)

type meanDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// Returns a new mean deviation indicator
func NewMeanDeviationIndicator(indicator Indicator, window int) Indicator {
	return meanDeviationIndicator{
		Indicator:     indicator,
		movingAverage: NewSimpleMovingAverage(indicator, window),
		window:        window,
	}
}

func (mdi meanDeviationIndicator) Calculate(index int) float64 {
	absoluteDeviations := 0.0

	average := mdi.movingAverage.Calculate(index)
	start := Max(0, index-mdi.window+1)

	for i := start; i <= index; i++ {
		absoluteDeviations = absoluteDeviations + math.Abs(mdi.Indicator.Calculate(i)-average)
	}

	return absoluteDeviations / float64(index-start+1)
}
