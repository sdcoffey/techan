package talib4g

type Strategy interface {
	ShouldEnter(index int, record *TradingRecord) bool
	ShouldExit(index int, record *TradingRecord) bool
}

type RuleStrategy struct {
	EntryRule      Rule
	ExitRule       Rule
	UnstablePeriod int
}

func (this RuleStrategy) ShouldEnter(index int, record *TradingRecord) bool {
	if this.EntryRule == nil {
		panic("entryrule is nil")
	}
	if index > this.UnstablePeriod && record.CurrentTrade().IsNew() {
		return this.EntryRule.IsSatisfied(index, record)
	}

	return false
}

func (this RuleStrategy) ShouldExit(index int, record *TradingRecord) bool {
	if index > this.UnstablePeriod && record.CurrentTrade().IsOpen() {
		return this.ExitRule.IsSatisfied(index, record)
	}

	return false
}
