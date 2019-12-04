package techan

import "github.com/sdcoffey/big"

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given window.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	return smaIndicator{indicator, window}
}

func (sma smaIndicator) Calculate(index int) big.Decimal {
	sum := big.ZERO
	for i := Max(0, index-sma.window+1); i <= index; i++ {
		sum = sum.Add(sma.indicator.Calculate(i))
	}
	realwindow := Min(sma.window, index+1)

	return sum.Div(big.NewDecimal(float64(realwindow)))
}

type emaIndicator struct {
	Indicator
	window      int
	alpha       big.Decimal
	resultCache []*big.Decimal
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given window, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	return &emaIndicator{
		Indicator:   indicator,
		window:      window,
		alpha:       big.NewDecimal(2.0 / float64(window+1)),
		resultCache: make([]*big.Decimal, 10000),
	}
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return ema.Indicator.Calculate(index)
	} else if len(ema.resultCache) > index && ema.resultCache[index] != nil {
		return *ema.resultCache[index]
	}

	emaPrev := ema.Calculate(index - 1)
	result := ema.Indicator.Calculate(index).Sub(emaPrev).Mul(ema.alpha).Add(emaPrev)
	ema.cacheResult(index, result)

	return result
}

func (ema *emaIndicator) cacheResult(index int, val big.Decimal) {
	if index < len(ema.resultCache) {
		ema.resultCache[index] = &val
	} else {
		ema.resultCache = append(ema.resultCache, &val)
	}
}

// NewMMAIndicator returns a derivative indciator which returns the modified moving average of the underlying
// indictator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	return &emaIndicator{
		Indicator:   indicator,
		window:      window,
		alpha:       big.NewDecimal(1.0 / float64(window)),
		resultCache: make([]*big.Decimal, 10000),
	}
}

// NewMACDIndicator returns a derivative Indicator which returns the difference between two EMAIndicators with long and
// short windows. It's useful for gauging the strength of price movements. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/m/macd.asp
func NewMACDIndicator(baseIndicator Indicator, shortwindow, longwindow int) Indicator {
	return NewDifferenceIndicator(NewEMAIndicator(baseIndicator, shortwindow), NewEMAIndicator(baseIndicator, longwindow))
}

// NewMACDHistogramIndicator returns a derivative Indicator based on the MACDIndicator, the result of which is
// the macd indicator minus it's signalLinewindow EMA. A more in-depth explanation can be found here:
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:macd-histogram
func NewMACDHistogramIndicator(macdIdicator Indicator, signalLinewindow int) Indicator {
	return NewDifferenceIndicator(macdIdicator, NewEMAIndicator(macdIdicator, signalLinewindow))
}
