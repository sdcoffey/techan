package talib4g

type positionNewRule struct{}

func (pnr positionNewRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsNew()
}

// NewPositionNewRule returns a new Rule that is satisfied when the current position in the trading record is new (no
// open positions).
func NewPositionNewRule() Rule {
	return positionNewRule{}
}

type positionOpenRule struct{}

// NewPositionOpenRule returns a new Rule that is satisfied when the current position in the trading record is open (position
// has been entered but not exited).
func NewPositionOpenRule() Rule {
	return positionOpenRule{}
}

func (pnr positionOpenRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsOpen()
}
