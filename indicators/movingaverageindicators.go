package indicators

import (
	"github.com/shopspring/decimal"
)

type SimpleMovingAverageIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this SimpleMovingAverageIndicator) Calculate(index int) decimal.Decimal {
	sum := decimal.Zero
	for i := max(0, index-this.TimeFrame+1); i <= index; i++ {
		sum = sum.Add(this.Indicator.Calculate(i))
	}
	realTimeFrame := min(this.TimeFrame, index+1)

	return sum.Div(decimal.New(realTimeFrame, 0))
}

type ExponentialMovingAverageIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this ExponentialMovingAverageIndicator) Calculate(index int) decimal.Decimal {
	if index+1 < this.TimeFrame {
		return SimpleMovingAverageIndicator{this.Indicator, this.TimeFrame}.Calculate(index)
	} else if index == 0 {
		return this.Indicator.Calculate(index)
	}

	emaPrev := this.Calculate(index - 1)
	multiplier := (decimal.New(1, 0) * 2).Div(decimal.New(this.TimeFrame+1, 0))
	return this.Indicator.Calculate(index).Sub(emaPrev).Mul(multiplier).Add(emaPrev)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
