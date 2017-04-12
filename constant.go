package talib4g

type ConstantIndicator float64

func (this ConstantIndicator) Calculate(index int) float64 {
	return float64(this)
}

type OverIndicator struct {
	High Indicator
	Low  Indicator
}
