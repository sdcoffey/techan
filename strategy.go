package talib4g

type Strategy struct {
	EntryRule      Rule
	ExitRule       Rule
	UnstablePeriod int
}

func (this Strategy) ShouldEnter(index int, record *TradingRecord) bool {
	if index > this.UnstablePeriod && record.CurrentTrade().IsNew() {
		return this.EntryRule.IsSatisfied(index, record)
	}

	return false
}

func (this Strategy) ShouldExit(index int, record *TradingRecord) bool {
	if index > this.UnstablePeriod && record.CurrentTrade().IsOpen() {
		return this.ExitRule.IsSatisfied(index, record)
	}

	return false
}
