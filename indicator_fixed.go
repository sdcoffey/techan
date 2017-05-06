package talib4g

type fixedIndicator []float64

func NewFixedIndicator(vals ...float64) Indicator {
	slc := make([]float64, len(vals))
	for i, val := range vals {
		slc[i] = val
	}

	return fixedIndicator(slc)
}

func (fi fixedIndicator) Calculate(index int) float64 {
	return fi[index]
}
