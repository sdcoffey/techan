package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type aroonIndicator struct {
	indicator Indicator
	window    int
	direction big.Decimal
	lowIndex  int
}

func (ai *aroonIndicator) Calculate(index int) big.Decimal {
	if index < ai.window-1 {
		return big.ZERO
	}

	oneHundred := big.TEN.Mul(big.TEN)
	ai.lowIndex = ai.findLowIndex(index)
	pSince := big.NewDecimal(float64(index - ai.lowIndex))
	windowAsDecimal := big.NewDecimal(float64(ai.window))

	return windowAsDecimal.Sub(pSince).Div(windowAsDecimal).Mul(oneHundred)
}

func (ai aroonIndicator) findLowIndex(index int) int {
	if ai.lowIndex < 1 || ai.lowIndex < index-ai.window {
		lv := big.NewDecimal(math.MaxFloat64)
		lowIndex := -1
		for i := (index + 1) - ai.window; i <= index; i++ {
			value := ai.indicator.Calculate(i).Mul(ai.direction)
			if value.LT(lv) {
				lv = value
				lowIndex = i
			}
		}

		return lowIndex
	}

	v1 := ai.indicator.Calculate(index).Mul(ai.direction)
	v2 := ai.indicator.Calculate(ai.lowIndex).Mul(ai.direction)

	if v1.LT(v2) {
		return index
	}

	return ai.lowIndex
}

// NewAroonUpIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the highest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a HighPriceIndicator or a derivative thereof
func NewAroonUpIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: big.ONE.Neg(),
		lowIndex:  -1,
	}
}

// NewAroonDownIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the lowest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a LowPriceIndicator or a derivative thereof
func NewAroonDownIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: big.ONE,
		lowIndex:  -1,
	}
}
