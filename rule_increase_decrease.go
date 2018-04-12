package techan

// IncreaseRule is satisfied when the given Indicator at the given index is greater than the value at the previous
// index.
type IncreaseRule struct {
	Indicator
}

// IsSatisfied returns true when the given Indicator at the given index is greater than the value at the previous
// index.
func (ir IncreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	if index == 0 {
		return false
	}

	return ir.Calculate(index).GT(ir.Calculate(index - 1))
}

// DecreaseRule is satisfied when the given Indicator at the given index is less than the value at the previous
// index.
type DecreaseRule struct {
	Indicator
}

// IsSatisfied returns true when the given Indicator at the given index is less than the value at the previous
// index.
func (dr DecreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	if index == 0 {
		return false
	}

	return dr.Calculate(index).LT(dr.Calculate(index - 1))
}
