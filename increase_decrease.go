package talib4g

type IncreaseRule struct {
	Indicator
}

func (ir IncreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	if index == 0 {
		return false
	} else {
		return ir.Calculate(index) > ir.Calculate(index-1)
	}
}

type DecreaseRule struct {
	Indicator
}

func (dr DecreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	if index == 0 {
		return false
	} else {
		return dr.Calculate(index) < dr.Calculate(index-1)
	}
}
