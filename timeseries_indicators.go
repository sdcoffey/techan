package talib4g

type AmountIndictor struct {
	*TimeSeries
}

type pl func(i int) float64

func (this AmountIndictor) Calculate(index int) float64 {
	return this.Ticks[index].Amount
}

type VolumeIndicator struct {
	*TimeSeries
}

func (this VolumeIndicator) Calculate(index int) float64 {
	return this.Ticks[index].Volume
}

type ClosePriceIndicator struct {
	*TimeSeries
}

func (this ClosePriceIndicator) Calculate(index int) float64 {
	return this.Ticks[index].ClosePrice
}

type TypicalPriceIndicator struct {
	*TimeSeries
}

func (this TypicalPriceIndicator) Calculate(index int) float64 {
	return (this.Ticks[index].MaxPrice + this.Ticks[index].MinPrice + this.Ticks[index].ClosePrice) / 3.0
}
