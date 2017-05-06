package talib4g

type constantIndicator float64

func NewConstantIndicator(constant float64) Indicator {
	return constantIndicator(constant)
}

func (ci constantIndicator) Calculate(index int) float64 {
	return float64(ci)
}
