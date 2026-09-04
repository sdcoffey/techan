package techan

import "github.com/sdcoffey/big"

type windowedStandardDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewWindowedStandardDeviationIndicator returns a indicator which calculates the standard deviation of the underlying
// indicator over a window. It returns zero before a complete window of usable
// input observations exists.
func NewWindowedStandardDeviationIndicator(ind Indicator, window int) Indicator {
	return windowedStandardDeviationIndicator{
		Indicator:     ind,
		movingAverage: NewSimpleMovingAverage(ind, window),
		window:        window,
	}
}

func (sdi windowedStandardDeviationIndicator) Calculate(index int) big.Decimal {
	if index < sdi.FirstValidIndex() {
		return big.ZERO
	}
	avg := sdi.movingAverage.Calculate(index)
	variance := big.ZERO
	for i := Max(0, index-sdi.window+1); i <= index; i++ {
		pow := sdi.Indicator.Calculate(i).Sub(avg).Pow(2)
		variance = variance.Add(pow)
	}
	realwindow := Min(sdi.window, index+1)

	return variance.Div(big.NewDecimal(float64(realwindow))).Sqrt()
}
