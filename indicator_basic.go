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

type highPriceIndicator struct {
	*TimeSeries
}

func NewHighPriceIndicator(series *TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) big.Decimal {
	return hpi.Candles[index].MaxPrice
}

type lowPriceIndicator struct {
	*TimeSeries
}

func NewLowPriceIndicator(series *TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) big.Decimal {
	return lpi.Candles[index].MinPrice
}

type openPriceIndicator struct {
	*TimeSeries
}

func NewOpenPriceIndicator(series *TimeSeries) Indicator {
	return openPriceIndicator{
		series,
	}
}

func (opi openPriceIndicator) Calculate(index int) big.Decimal {
	return opi.Candles[index].OpenPrice
}

type typicalPriceIndicator struct {
	*TimeSeries
}

func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) big.Decimal {
	numerator := tpi.Candles[index].MaxPrice.Add(tpi.Candles[index].MinPrice).Add(tpi.Candles[index].ClosePrice)
	return numerator.Div(big.NewFromString("3"))
}
