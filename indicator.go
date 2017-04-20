package talib4g

import "time"

type Indicator interface {
	Calculate(int) float64
}

func Plot(xvals []time.Time, i Indicator) []float64 {
	y := make([]float64, len(xvals))
	for idx := range xvals {
		y[idx] = i.Calculate(idx)
	}

	return y
}
