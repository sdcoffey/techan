package techan

import "github.com/sdcoffey/big"

// NewMaximumValueIndicator returns a derivative Indicator which returns the maximum value
// present in a given window. Use a window value of -1 to include all values in the
// underlying indicator.
func NewMaximumValueIndicator(ind Indicator, window int) Indicator {
	return maximumValueIndicator{
		indicator: ind,
		window:    window,
	}
}

type maximumValueIndicator struct {
	indicator Indicator
	window    int
}

func (mvi maximumValueIndicator) Calculate(index int) big.Decimal {
	maxValue := big.NewFromString("-Inf")

	start := 0
	if mvi.window > 0 {
		start = Max(index-mvi.window+1, 0)
	}

	for i := start; i <= index; i++ {
		value := mvi.indicator.Calculate(i)
		if value.GT(maxValue) {
			maxValue = value
		}
	}

	return maxValue
}
