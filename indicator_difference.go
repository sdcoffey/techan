package talib4g

import "github.com/sdcoffey/big"

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

func (di differenceIndicator) Calculate(index int) big.Decimal {
	return di.minuend.Calculate(index).Sub(di.subtrahend.Calculate(index))
}
