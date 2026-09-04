package techan

import (
	"github.com/sdcoffey/big"
)

type aroonIndicator struct {
	indicator Indicator
	window    int
	direction big.Decimal
}

func (ai *aroonIndicator) Calculate(index int) big.Decimal {
	if index < ai.FirstValidIndex() {
		return big.ZERO
	}
	start := index - ai.window + 1
	extremeIndex := start
	extreme := ai.indicator.Calculate(start).Mul(ai.direction)
	for i := start; i <= index; i++ {
		value := ai.indicator.Calculate(i).Mul(ai.direction)
		if value.NaN() {
			return big.NaN
		}
		// The latest occurrence of an equal extreme determines periods since it.
		if value.LTE(extreme) {
			extreme, extremeIndex = value, i
		}
	}
	return big.NewFromInt(ai.window - (index - extremeIndex)).Div(big.NewFromInt(ai.window)).Mul(big.NewFromInt(100))
}

// NewAroonUpIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the highest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a HighPriceIndicator or a derivative thereof
func NewAroonUpIndicator(indicator Indicator, window int) Indicator {
	requirePositiveWindow(window)
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: big.ONE.Neg(),
	}
}

// NewAroonDownIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the lowest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a LowPriceIndicator or a derivative thereof
func NewAroonDownIndicator(indicator Indicator, window int) Indicator {
	requirePositiveWindow(window)
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: big.ONE,
	}
}
