package talib4g

type differenceIndicator struct {
	minuend    Indicator
	subtrahend Indicator
}

func NewDifferenceIndicator(minuend, subtrahend Indicator) Indicator {
	return differenceIndicator{
		minuend:    minuend,
		subtrahend: subtrahend,
	}
}

func (di differenceIndicator) Calculate(index int) Decimal {
	return di.minuend.Calculate(index).Sub(di.subtrahend.Calculate(index))
}
