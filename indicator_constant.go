package talib4g

type constantIndicator float64

func NewConstantIndicator(constant float64) Indicator {
	return constantIndicator(constant)
}

func (ci constantIndicator) Calculate(index int) Decimal {
	return NewDecimal(float64(ci))
}
