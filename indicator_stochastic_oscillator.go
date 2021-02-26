package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type kIndicator struct {
	closePrice Indicator
	minValue   Indicator
	maxValue   Indicator
	window     int
}

// NewFastStochasticIndicator returns a derivative Indicator which returns the fast stochastic indicator (%K) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewFastStochasticIndicator(closePrice, minValue, maxValue Indicator, window int) Indicator {
	return kIndicator{closePrice, minValue, maxValue, window}
}

func (k kIndicator) Calculate(index int) big.Decimal {
	closeVal := k.closePrice.Calculate(index)
	minVal := k.minValue.Calculate(index)
	maxVal := k.maxValue.Calculate(index)

	if minVal.EQ(maxVal) {
		return big.NewDecimal(math.MaxFloat64)
	}

	return closeVal.Sub(minVal).Div(maxVal.Sub(minVal)).Mul(big.NewDecimal(100))
}

type dIndicator struct {
	k      Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	return dIndicator{k, window}
}

func (d dIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(d.k, d.window).Calculate(index)
}
