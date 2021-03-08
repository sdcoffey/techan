package techan

import (
	"github.com/sdcoffey/big"
)

type stochasticRSIIndicator struct {
	curRSI Indicator
	minRSI Indicator
	maxRSI Indicator
}

// NewStochasticRSIIndicator returns a derivative Indicator which returns the stochastic RSI indicator for the given
// RSI window.
// https://www.investopedia.com/terms/s/stochrsi.asp
func NewStochasticRSIIndicator(indicator Indicator, timeframe int) Indicator {
	rsiIndicator := NewRelativeStrengthIndexIndicator(indicator, timeframe)
	return stochasticRSIIndicator{
		curRSI: rsiIndicator,
		minRSI: NewMinimumValueIndicator(rsiIndicator, timeframe),
		maxRSI: NewMaximumValueIndicator(rsiIndicator, timeframe),
	}
}

func (sri stochasticRSIIndicator) Calculate(index int) big.Decimal {
	curRSI := sri.curRSI.Calculate(index)
	minRSI := sri.minRSI.Calculate(index)
	maxRSI := sri.maxRSI.Calculate(index)

	if minRSI.EQ(maxRSI) {
		return big.NewDecimal(100)
	}

	return curRSI.Sub(minRSI).Div(maxRSI.Sub(minRSI)).Mul(big.NewDecimal(100))
}

type stochRSIKIndicator struct {
	stochasticRSI Indicator
	window        int
}

// NewFastStochasticRSIIndicator returns a derivative Indicator which returns the fast stochastic RSI indicator (%K)
// for the given stochastic window.
func NewFastStochasticRSIIndicator(stochasticRSI Indicator, timeframe int) Indicator {
	return stochRSIKIndicator{stochasticRSI, timeframe}
}

func (k stochRSIKIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(k.stochasticRSI, k.window).Calculate(index)
}

type stochRSIDIndicator struct {
	fastStochasticRSI Indicator
	window            int
}

// NewSlowStochasticRSIIndicator returns a derivative Indicator which returns the slow stochastic RSI indicator (%D)
// for the given stochastic window.
func NewSlowStochasticRSIIndicator(fastStochasticRSI Indicator, timeframe int) Indicator {
	return stochRSIDIndicator{fastStochasticRSI, timeframe}
}

func (d stochRSIDIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(d.fastStochasticRSI, d.window).Calculate(index)
}
