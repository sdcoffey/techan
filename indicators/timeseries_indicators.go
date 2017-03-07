package indicators

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
)

type AmountIndictor *TimeSeries

func (this AmountIndictor) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].Amount
}

type VolumeIndicator *TimeSeries

func (this VolumeIndicator) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].Volume
}

type ClosePriceIndicator *TimeSeries

func (this ClosePriceIndicator) Calculate(index int) decimal.Decimal {
	return this.Ticks[index].ClosePrice
}
