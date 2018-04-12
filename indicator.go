package techan

import "github.com/sdcoffey/big"

// Indicator is an interface that describes a methodology by which to analyze a trading record for a specific property
// or trend. For example. MovingAverageIndicator implements the Indicator interface and, for a given index in the timeSeries,
// returns the current moving average of the prices in that series.
type Indicator interface {
	Calculate(int) big.Decimal
}
