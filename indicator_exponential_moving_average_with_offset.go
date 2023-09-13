package techan

import "github.com/sdcoffey/big"

type emaIndicatorWithOffset struct {
	indicator   Indicator
	window      int
	offset      int
	alpha       big.Decimal
	resultCache resultCache
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicatorWithOffset(indicator Indicator, window int, offset int) Indicator {
	return &emaIndicatorWithOffset{
		indicator:   indicator,
		window:      window,
		offset:      offset,
		alpha:       big.ONE.Frac(2).Div(big.NewFromInt(window + 1)),
		resultCache: make([]*big.Decimal, 1000),
	}
}

func (ema *emaIndicatorWithOffset) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCachedWithOffset(ema, index, ema.offset, func(i int) big.Decimal {
		return NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := ema.indicator.Calculate(index).Mul(ema.alpha)
	result := todayVal.Add(ema.Calculate(index - 1).Mul(big.ONE.Sub(ema.alpha)))

	cacheResult(ema, index, result)

	return result
}

func (ema emaIndicatorWithOffset) cache() resultCache { return ema.resultCache }

func (ema *emaIndicatorWithOffset) setCache(newCache resultCache) {
	ema.resultCache = newCache
}

func (ema emaIndicatorWithOffset) windowSize() int { return ema.window }
