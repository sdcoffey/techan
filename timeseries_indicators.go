package talib4g

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
