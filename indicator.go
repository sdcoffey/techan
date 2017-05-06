package talib4g

type Indicator interface {
	Calculate(int) float64
}
