package techan

import "github.com/sdcoffey/big"

// Indicator is an interface that describes a methodology by which to analyze a trading record for a specific property
// or trend. For example. MovingAverageIndicator implements the Indicator interface and, for a given index in the timeSeries,
// returns the current moving average of the prices in that series.
type Indicator interface {
	Calculate(int) big.Decimal
}

// WarmupIndicator reports the first index with enough observations to calculate
// a value. Earlier values are placeholders and must not seed another indicator.
type WarmupIndicator interface {
	FirstValidIndex() int
}

// FirstValidIndex returns an indicator's first usable index. Custom indicators
// without warm-up metadata are assumed to be usable from index zero.
func FirstValidIndex(indicator Indicator) int {
	if ready, ok := indicator.(WarmupIndicator); ok {
		return max(0, ready.FirstValidIndex())
	}
	return 0
}

func requirePositiveWindow(window int) {
	if window <= 0 {
		panic("indicator window must be positive")
	}
}
