package indicators

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
)

type CCIIndicator struct {
	tpi TypicalPriceIndicator
	sma SMAIndicator
	md  MeanDeviationIndicator
	tf  int
}

func NewCCIIndicator(ts *TimeSeries, timeFrame int) CCIIndicator {
	return CCIIndicator{
		tpi: TypicalPriceIndicator{ts},
		sma: SMAIndicator{TypicalPriceIndicator{ts}, timeFrame},
		md:  NewMeanDeviationIndicator(TypicalPriceIndicator{ts}, timeFrame),
		tf:  timeFrame,
	}
}

func (this CCIIndicator) Calculate(index int) decimal.Decimal {
	typicalPrice := this.tpi.Calculate(index)
	typicalPriceAvg := this.sma.Calculate(index)
	mean := this.md.Calculate(index)
	if mean.Cmp(ZERO) == 0 {
		return ZERO
	}

	return (typicalPrice.Sub(typicalPriceAvg)).Div(mean.Mul(decimal.NewFromFloat(0.015)))
}
