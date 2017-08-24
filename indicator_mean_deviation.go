package talib4g

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

func (mdi meanDeviationIndicator) Calculate(index int) Decimal {
	average := mdi.movingAverage.Calculate(index)
	start := Max(0, index-mdi.window+1)
	absoluteDeviations := NewDecimal(0)

	for i := start; i <= index; i++ {
		absoluteDeviations = absoluteDeviations.Add(average.Sub(mdi.Indicator.Calculate(i)).Abs())
	}

	return absoluteDeviations.Div(NewDecimal(float64(Min(mdi.window, index-start+1))))
}
