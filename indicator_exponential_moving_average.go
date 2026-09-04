package techan

import (
	mathbig "math/big"

	"github.com/sdcoffey/big"
)

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       big.Decimal
	resultCache resultCache
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. It returns zero for indices before
// the first complete window. A more in-depth explanation can be found here: http://www.investopedia.com/terms/e/ema.asp
// Cached values are rounded to 256 binary bits (approximately 77 significant
// decimal digits) to bound memory usage over arbitrarily long histories.
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	requirePositiveWindow(window)
	return &emaIndicator{
		indicator: indicator,
		window:    window,
		alpha:     big.NewFromInt(2).Div(big.NewFromInt(window).Add(big.ONE)),
	}
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	first := ema.FirstValidIndex()
	if index < first {
		return big.ZERO
	}
	if index < len(ema.resultCache) && ema.resultCache[index] != nil {
		return *ema.resultCache[index]
	}
	if index >= len(ema.resultCache) {
		expandResultCache(ema, index+1)
	}
	if ema.window == 1 {
		cacheResult(ema, index, roundSmoothingValue(ema.indicator.Calculate(index)))
		return *ema.resultCache[index]
	}

	// Fill forward from the closest cached predecessor, without recursion.
	start := index
	for start > first && ema.resultCache[start-1] == nil {
		start--
	}
	if start == first {
		seed := NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(first)
		cacheResult(ema, first, roundSmoothingValue(seed))
		start++
	}
	weight := big.ONE.Sub(ema.alpha)
	for i := start; i <= index; i++ {
		value := ema.indicator.Calculate(i).Mul(ema.alpha).
			Add(ema.resultCache[i-1].Mul(weight))
		cacheResult(ema, i, roundSmoothingValue(value))
	}
	return *ema.resultCache[index]
}

func (ema emaIndicator) cache() resultCache { return ema.resultCache }

func (ema *emaIndicator) setCache(newCache resultCache) {
	ema.resultCache = newCache
}

func (ema emaIndicator) FirstValidIndex() int {
	return FirstValidIndex(ema.indicator) + ema.window - 1
}

// ResetCacheFrom invalidates cached EMA values from index onward.
func (ema *emaIndicator) ResetCacheFrom(index int) {
	resetResultCache(ema, index)
	ResetCacheFrom(ema.indicator, index)
}

// big v0.8 grows precision with each operation. Round recursive state without
// narrowing the exponent range or using float64. FormattedString(-1) preserves
// all digits independently of MarshalQuoted; String retains only ten digits.
func roundSmoothingValue(value big.Decimal) big.Decimal {
	if value.NaN() {
		return value
	}
	rounded, _, err := mathbig.ParseFloat(value.FormattedString(-1), 10, 256, mathbig.ToNearestEven)
	if err != nil {
		panic(err) // Decimal always formats a valid number or infinity.
	}
	return big.NewFromString(rounded.Text('g', -1))
}
