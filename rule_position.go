package techan

// PositionNewRule is satisfied when the current position in the trading record is new (no
// open positions).
type PositionNewRule struct{}

// IsSatisfied returns true if the current position in the record is new
func (pnr PositionNewRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsNew()
}

// PositionOpenRule is satisfied when the current position in the trading record is open (position
// has been entered but not exited).
type PositionOpenRule struct{}

// IsSatisfied returns true if the current position in the record is Open
func (pnr PositionOpenRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsOpen()
}
