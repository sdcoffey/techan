package talib4g

type CumulativeGainsIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this CumulativeGainsIndicator) Calculate(index int) float64 {
	result := 0.0
	for i := Max(1, index-this.TimeFrame+1); i <= index; i++ {
		if this.Indicator.Calculate(i) > this.Indicator.Calculate(i-1) {
			result += this.Indicator.Calculate(i) - this.Indicator.Calculate(i-1)
		}
	}

	return result
}

type CumulativeLossesIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this CumulativeLossesIndicator) Calculate(index int) float64 {
	result := 0.0
	for i := Max(1, index-this.TimeFrame+1); i <= index; i++ {
		if this.Indicator.Calculate(i) < this.Indicator.Calculate(i-1) {
			result += this.Indicator.Calculate(i) - this.Indicator.Calculate(i-1)
		}
	}

	return result
}

type AverageIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this AverageIndicator) Calculate(index int) float64 {
	return this.Indicator.Calculate(index) / float64(Min(index+1, this.TimeFrame))
}
