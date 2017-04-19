package talib4g

type Indicator interface {
	Calculate(int) float64
	//Plot(xvals []time.Time) chart.TimeSeries
}
