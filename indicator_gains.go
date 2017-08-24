package talib4g

type cumulativeIndicator struct {
	Indicator
	window int
	mult   Decimal
}

func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      ONE,
	}
}

func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      ONE.Neg(),
	}
}

func (ci cumulativeIndicator) Calculate(index int) Decimal {
	total := NewDecimal(0.0)

	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := ci.Indicator.Calculate(i).Sub(ci.Indicator.Calculate(i - 1))
		if diff.Mul(ci.mult).GT(ZERO) {
			total = total.Add(diff.Abs())
		}
	}

	return total
}

type percentChangeIndicator struct {
	Indicator
}

func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) Decimal {
	if index == 0 {
		return ZERO
	}

	cp := pgi.Indicator.Calculate(index)
	cplast := pgi.Indicator.Calculate(index - 1)
	return cp.Div(cplast).Sub(ONE)
}
