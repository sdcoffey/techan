package techan

import "github.com/sdcoffey/big"

// NewMMAIndicator returns Wilder's modified moving average. It seeds with the
// first full window of usable input values, then uses a weight of 1/window.
// Cached values use the same 256-bit precision policy as EMA.
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	requirePositiveWindow(window)
	return &emaIndicator{
		indicator: indicator,
		window:    window,
		alpha:     big.ONE.Div(big.NewFromInt(window)),
	}
}
