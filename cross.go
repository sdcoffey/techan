package talib4g

type CrossIndicator struct {
	Upper Indicator
	Lower Indicator
}

func (this CrossIndicator) Calculate(index int) float64 {
	i := index

	if i == 0 || this.Upper.Calculate(index) >= this.Lower.Calculate(i) {
		return 0
	}

	i--
	if this.Upper.Calculate(i) > this.Lower.Calculate(i) {
		return 1
	} else {
		for i > 0 && this.Upper.Calculate(i) == this.Lower.Calculate(i) {
			i--
		}
		if i != 0 && this.Upper.Calculate(i) > this.Lower.Calculate(i) {
			return 1
		}
	}

	return 0
}
