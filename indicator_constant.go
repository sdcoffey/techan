package talib4g

import "github.com/sdcoffey/big"

type constantIndicator float64

func NewConstantIndicator(constant float64) Indicator {
	return constantIndicator(constant)
}

func (ci constantIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(float64(ci))
}
