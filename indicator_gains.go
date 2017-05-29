package talib4g

import "math"

type cumulativeIndicator struct {
	Indicator
	window int
	mult   int
}

func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      1,
	}
}

func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      -1,
	}
}

func (ci cumulativeIndicator) Calculate(index int) float64 {
	total := 0.0

	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := ci.Indicator.Calculate(i) - ci.Indicator.Calculate(i-1)
		if diff*float64(ci.mult) > 0 {
			total += math.Abs(diff)
		}
	}

	return total
}

type percentChangeIndicator struct {
	Indicator
}

func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) float64 {
	if index == 0 {
		return 0
	}

	cp := pgi.Indicator.Calculate(index)
	cplast := pgi.Indicator.Calculate(index - 1)
	return cp/cplast - 1
	//return pgi.Indicator.Calculate(index)/pgi.Indicator.Calculate(index-1) - 1
}
