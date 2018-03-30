package talib4g

type positionNewRule struct{}

func (pnr positionNewRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsNew()
}

func NewPositionNewRule() Rule {
	return positionNewRule{}
}

type positionOpenRule struct{}

func NewPositionOpenRule() Rule {
	return positionOpenRule{}
}

func (pnr positionOpenRule) IsSatisfied(index int, record *TradingRecord) bool {
	return record.CurrentPosition().IsOpen()
}
