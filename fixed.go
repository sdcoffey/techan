package talib4g

type FixedIndicator []float64

func NewFixedIndicator(vals ...int) FixedIndicator {
	slc := make([]float64, len(vals))
	for i, val := range vals {
		slc[i] = float64(val)
	}

	return FixedIndicator(slc)
}

func (this FixedIndicator) Calculate(index int) float64 {
	return this[index]
}
