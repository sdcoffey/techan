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

func (di differenceIndicator) Calculate(index int) float64 {
	return di.minuend.Calculate(index) - di.subtrahend.Calculate(index)
}
