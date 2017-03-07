package talib4g

type Strategy struct {
	EntryRule      Rule
	ExitRule       Rule
	UnstablePeriod int
}

func (this Strategy) ShouldEnter(index int, record TradingRecord) bool {
	if index < this.UnstablePeriod {
		return false
	}

	return this.EntryRule.IsSatisfied(index, record)
}

func (this Strategy) ShouldExit(index int, record TradingRecord) bool {
	if index < this.UnstablePeriod {
		return false
	}

	return this.ExitRule.IsSatisfied(index, record)
}
