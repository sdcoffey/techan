package talib4g

type volumeIndicator struct {
	*TimeSeries
}

func NewVolumeIndicator(series *TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) float64 {
	return vi.Candles[index].Volume
}

type closePriceIndicator struct {
	*TimeSeries
}

func NewClosePriceIndicator(series *TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) float64 {
	return cpi.Candles[index].ClosePrice
}

type typicalPriceIndicator struct {
	*TimeSeries
}

func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (this typicalPriceIndicator) Calculate(index int) float64 {
	return (this.Candles[index].MaxPrice + this.Candles[index].MinPrice + this.Candles[index].ClosePrice) / 3.0
}
