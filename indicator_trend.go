package techan

import "github.com/sdcoffey/big"

type trendLineIndicator struct {
	indicator Indicator
	window    int
}

// NewTrendlineIndicator returns an indicator whose output is the slope of the trend
// line given by the values in the window.
func NewTrendlineIndicator(indicator Indicator, window int) Indicator {
	return trendLineIndicator{
		indicator: indicator,
		window:    window,
	}
}

func (tli trendLineIndicator) Calculate(index int) big.Decimal {
	window := Min(index+1, tli.window)

	values := make([]big.Decimal, window)

	for i := 0; i < window; i++ {
		values[i] = tli.indicator.Calculate(index - (window - 1) + i)
	}

	n := big.ONE.Mul(big.NewDecimal(float64(window)))
	ab := sumXy(values).Mul(n).Sub(sumX(values).Mul(sumY(values)))
	cd := sumX2(values).Mul(n).Sub(sumX(values).Pow(2))

	return ab.Div(cd)
}

func sumX(decimals []big.Decimal) (s big.Decimal) {
	s = big.ZERO

	for i := range decimals {
		s = s.Add(big.NewDecimal(float64(i)))
	}

	return s
}

func sumY(decimals []big.Decimal) (b big.Decimal) {
	b = big.ZERO
	for _, d := range decimals {
		b = b.Add(d)
	}

	return
}

func sumXy(decimals []big.Decimal) (b big.Decimal) {
	b = big.ZERO

	for i, d := range decimals {
		b = b.Add(d.Mul(big.NewDecimal(float64(i))))
	}

	return
}

func sumX2(decimals []big.Decimal) big.Decimal {
	b := big.ZERO

	for i := range decimals {
		b = b.Add(big.NewDecimal(float64(i)).Pow(2))
	}

	return b
}
