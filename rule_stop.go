package talib4g

import "github.com/sdcoffey/big"

type StopLossRule struct {
	Indicator
	tolerance big.Decimal
}

// Returns a new stop loss rule based on a timeseries and a loss tolerance
// The loss tolerance should be a number between -1 and 1, where negative
// values represent a loss and vice versa.
func NewStopLossRule(series *TimeSeries, lossTolerance float64) Rule {
	return StopLossRule{
		Indicator: NewClosePriceIndicator(series),
		tolerance: big.NewDecimal(lossTolerance),
	}
}

func (slr StopLossRule) IsSatisfied(index int, record *TradingRecord) bool {
	if !record.CurrentPosition().IsOpen() {
		return false
	}

	openPrice := record.CurrentPosition().CostBasis()
	loss := slr.Indicator.Calculate(index).Div(openPrice).Sub(big.ONE)
	return loss.LTE(slr.tolerance)
}
