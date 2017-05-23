package talib4g

import "math"

type Rule interface {
	IsSatisfied(index int, record *TradingRecord) bool
}

func And(r1, r2 Rule) Rule {
	return andRule{r1, r2}
}

func Or(r1, r2 Rule) Rule {
	return orRule{r1, r2}
}

type andRule struct {
	r1 Rule
	r2 Rule
}

func (this andRule) IsSatisfied(index int, record *TradingRecord) bool {
	return this.r1.IsSatisfied(index, record) && this.r2.IsSatisfied(index, record)
}

type orRule struct {
	r1 Rule
	r2 Rule
}

func (this orRule) IsSatisfied(index int, record *TradingRecord) bool {
	return this.r1.IsSatisfied(index, record) || this.r2.IsSatisfied(index, record)
}

type OverIndicatorRule struct {
	First  Indicator
	Second Indicator
}

func (this OverIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return this.First.Calculate(index) > this.Second.Calculate(index)
}

type UnderIndicatorRule struct {
	First  Indicator
	Second Indicator
}

func (this UnderIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return this.First.Calculate(index) < this.Second.Calculate(index)
}

type percentChangeRule struct {
	indicator Indicator
	percent   float64
}

func (pgr percentChangeRule) IsSatisfied(index int, record *TradingRecord) bool {
	return math.Abs(pgr.indicator.Calculate(index)) > math.Abs(pgr.percent)
}

func NewPercentChangeRule(indicator Indicator, percent float64) Rule {
	return percentChangeRule{
		indicator: NewPercentChangeIndicator(indicator),
		percent: percent,
	}
}
