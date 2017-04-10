package talib4g

import (
	"github.com/shopspring/decimal"
)

type AmountIndictor struct {
	*TimeSeries
}

func (this AmountIndictor) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].Amount
}

type VolumeIndicator struct {
	*TimeSeries
}

func (this VolumeIndicator) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].Volume
}

type ClosePriceIndicator struct {
	*TimeSeries
}

func (this ClosePriceIndicator) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].ClosePrice
}

type TypicalPriceIndicator struct {
	*TimeSeries
}

func (this TypicalPriceIndicator) Calculate(index int) decimal.Decimal {
	maxPrice := this.Ticks[index].MaxPrice
	minPrice := this.Ticks[index].MinPrice
	closePrice := this.Ticks[index].ClosePrice

	return maxPrice.Add(minPrice).Add(closePrice).Div(THREE)
}
