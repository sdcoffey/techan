package techan

import "github.com/sdcoffey/big"

// NewMinimumValueIndicator returns a derivative Indicator which returns the minimum value
// present in a given window. Use a window value of -1 to include all values in the
// underlying indicator.
func NewMinimumValueIndicator(ind Indicator, window int) Indicator {
	return minimumValueIndicator{
		indicator: ind,
		window:    window,
	}
}

type minimumValueIndicator struct {
	indicator Indicator
	window    int
}

func (mvi minimumValueIndicator) Calculate(index int) big.Decimal {
	minValue := big.NewFromString("Inf")

	start := mvi.FirstValidIndex()
	if index < start {
		return big.ZERO
	}
	if mvi.window > 0 {
		start = max(index-mvi.window+1, start)
	}

	for i := start; i <= index; i++ {
		value := mvi.indicator.Calculate(i)
		if value.LT(minValue) {
			minValue = value
		}
	}

	return minValue
}
