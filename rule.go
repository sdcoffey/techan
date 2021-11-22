package techan

import "github.com/sdcoffey/big"

// Rule is an interface describing an algorithm by which a set of criteria may be satisfied
type Rule interface {
	IsSatisfied(index int, record *TradingRecord) bool
}

// And returns a new rule whereby BOTH of the passed-in rules must be satisfied for the rule to be satisfied
func And(r ...Rule) Rule {
	return andRule{r}
}

// Or returns a new rule whereby ONE OF the passed-in rules must be satisfied for the rule to be satisfied
func Or(r ...Rule) Rule {
	return orRule{r}
}

type andRule struct {
	r []Rule
}

func (ar andRule) IsSatisfied(index int, record *TradingRecord) bool {
	for _, r := range ar.r {
		if !r.IsSatisfied(index, record) {
			return false
		}
	}
	return true
}

type orRule struct {
	r []Rule
}

func (or orRule) IsSatisfied(index int, record *TradingRecord) bool {
	for _, r := range or.r {
		if r.IsSatisfied(index, record) {
			return true
		}
	}
	return false
}

// OverIndicatorRule is a rule where the First Indicator must be greater than the Second Indicator to be Satisfied
type OverIndicatorRule struct {
	First  Indicator
	Second Indicator
}

// IsSatisfied returns true when the First Indicator is greater than the Second Indicator
func (oir OverIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return oir.First.Calculate(index).GT(oir.Second.Calculate(index))
}

// UnderIndicatorRule is a rule where the First Indicator must be less than the Second Indicator to be Satisfied
type UnderIndicatorRule struct {
	First  Indicator
	Second Indicator
}

// IsSatisfied returns true when the First Indicator is less than the Second Indicator
func (uir UnderIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	return uir.First.Calculate(index).LT(uir.Second.Calculate(index))
}

type percentChangeRule struct {
	indicator Indicator
	percent   big.Decimal
}

func (pgr percentChangeRule) IsSatisfied(index int, record *TradingRecord) bool {
	return pgr.indicator.Calculate(index).Abs().GT(pgr.percent.Abs())
}

// NewPercentChangeRule returns a rule whereby the given Indicator must have changed by a given percentage to be satisfied.
// You should specify percent as a float value between -1 and 1
func NewPercentChangeRule(indicator Indicator, percent float64) Rule {
	return percentChangeRule{
		indicator: NewPercentChangeIndicator(indicator),
		percent:   big.NewDecimal(percent),
	}
}
