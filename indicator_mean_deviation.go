package talib4g

import "github.com/sdcoffey/big"

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

func (mdi meanDeviationIndicator) Calculate(index int) big.Decimal {
	average := mdi.movingAverage.Calculate(index)
	start := Max(0, index-mdi.window+1)
	absoluteDeviations := big.NewDecimal(0)

	for i := start; i <= index; i++ {
		absoluteDeviations = absoluteDeviations.Add(average.Sub(mdi.Indicator.Calculate(i)).Abs())
	}

	return absoluteDeviations.Div(big.NewDecimal(float64(Min(mdi.window, index-start+1))))
}
