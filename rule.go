package talib4g

type Rule interface {
	IsSatisfied(index int, record *TradingRecord) bool
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

type ruleBase struct {
	Rule
}

func (this ruleBase) And(other Rule) Rule {
	return andRule{this, other}
}

func (this ruleBase) Or(other Rule) Rule {
	return orRule{this, other}
}

type SimpleRule struct {
	ruleBase
}

func (this SimpleRule) IsSatisfied(index int, record *TradingRecord) bool {
	return index > 0
}
