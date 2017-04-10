package talib4g

import (
	"github.com/shopspring/decimal"
)

type CumulativeGainsIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this CumulativeGainsIndicator) Calculate(index int) decimal.Decimal {
	result := ZERO
	for i := Max(1, index-this.TimeFrame+1); i <= index; i++ {
		if this.Indicator.Calculate(i).Cmp(this.Indicator.Calculate(i-1)) > 0 {
			result = result.Add(this.Indicator.Calculate(i).Sub(this.Indicator.Calculate(i - 1)))
		}
	}

	return result
}

type CumulativeLossesIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this CumulativeLossesIndicator) Calculate(index int) decimal.Decimal {
	result := ZERO
	for i := Max(1, index-this.TimeFrame+1); i <= index; i++ {
		if this.Indicator.Calculate(i).Cmp(this.Indicator.Calculate(i-1)) < 0 {
			result = result.Add(this.Indicator.Calculate(i).Sub(this.Indicator.Calculate(i - 1)))
		}
	}

	return result
}

type AverageIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this AverageIndicator) Calculate(index int) decimal.Decimal {
	return this.Indicator.Calculate(index).Div(NewDecimal(Min(index+1, this.TimeFrame)))
}
