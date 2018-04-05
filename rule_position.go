package talib4g

// PositionNewRule is satisfied when the current position in the trading record is new (no
// open positions).
type PositionNewRule struct{}

func (pnr PositionNewRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsNew()
}

//PositionOpenRule is satisfied when the current position in the trading record is open (position
// has been entered but not exited).
type PositionOpenRule struct{}

func (pnr PositionOpenRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsOpen()
}
