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

	start := 0
	if mvi.window > 0 {
		start = Max(index-mvi.window+1, 0)
	}

	for i := start; i <= index; i++ {
		value := mvi.indicator.Calculate(i)
		if value.LT(minValue) {
			minValue = value
		}
	}

	return minValue
}
