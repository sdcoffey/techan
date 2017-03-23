package indicators

import (
	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
)

type SMAIndicator struct {
	Indicator Indicator
	TimeFrame int
}

func (this SMAIndicator) Calculate(index int) decimal.Decimal {
	sum := ZERO
	for i := Max(0, index-this.TimeFrame+1); i <= index; i++ {
		sum = sum.Add(this.Indicator.Calculate(i))
	}
	realTimeFrame := Min(this.TimeFrame, index+1)

	return sum.Div(NewDecimal(realTimeFrame))
}

type EMAIndicator struct {
	Indicator   Indicator
	TimeFrame   int
	resultCache []*decimal.Decimal
}

func NewEMAIndicator(indicator Indicator, timeFrame int) *EMAIndicator {
	return &EMAIndicator{
		Indicator:   indicator,
		TimeFrame:   timeFrame,
		resultCache: make([]*decimal.Decimal, timeFrame),
	}
}

func (this *EMAIndicator) Calculate(index int) decimal.Decimal {
	if len(this.resultCache) > index && this.resultCache[index] != nil {
		return *this.resultCache[index]
	} else if index+1 < this.TimeFrame {
		result := SMAIndicator{this.Indicator, this.TimeFrame}.Calculate(index)
		this.cacheResult(index, result)

		return result
	} else if index == 0 {
		result := this.Indicator.Calculate(index)
		this.cacheResult(index, result)

		return result
	}

	emaPrev := this.Calculate(index - 1)
	result := this.Indicator.Calculate(index).Sub(emaPrev).Mul(this.multiplier()).Add(emaPrev)
	this.cacheResult(index, result)

	return result
}

func (this *EMAIndicator) cacheResult(index int, value decimal.Decimal) {
	if index < len(this.resultCache) {
		this.resultCache[index] = &value
	} else {
		this.resultCache = append(this.resultCache, &value)
	}
}

func (this EMAIndicator) multiplier() decimal.Decimal {
	return TWO.Div(NewDecimal(this.TimeFrame + 1))
}

type MACDIndicator struct {
	shortEMA Indicator
	longEMA  Indicator
}

func NewMACDIndicator(i Indicator, shortTimeFrame, longTimeFrame int) MACDIndicator {
	return MACDIndicator{
		NewEMAIndicator(i, shortTimeFrame),
		NewEMAIndicator(i, longTimeFrame),
	}
}

func (this MACDIndicator) Calculate(index int) decimal.Decimal {
	return this.shortEMA.Calculate(index).Sub(this.longEMA.Calculate(index))
}
