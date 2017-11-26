package talib4g

import "github.com/sdcoffey/big"

type volumeIndicator struct {
	*TimeSeries
}

func NewVolumeIndicator(series *TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) big.Decimal {
	return vi.Candles[index].Volume
}

type closePriceIndicator struct {
	*TimeSeries
}

func NewClosePriceIndicator(series *TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) big.Decimal {
	return cpi.Candles[index].ClosePrice
}

type typicalPriceIndicator struct {
	*TimeSeries
}

func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (this typicalPriceIndicator) Calculate(index int) big.Decimal {
	return this.Candles[index].MaxPrice.Add(this.Candles[index].MinPrice).Add(this.Candles[index].ClosePrice).Div(big.ONE.Frac(3))
}
