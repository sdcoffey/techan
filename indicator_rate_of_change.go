package techan

import "github.com/sdcoffey/big"

type rateOfChangeIndicator struct {
	series *TimeSeries
	window int
}

// NewRateOfChangeIndicator returns an indicator that calculates the Rate of Change (ROC)
// which measures the percentage change between the current price and the price n periods ago
// window: number of periods to look back
func NewRateOfChangeIndicator(series *TimeSeries, window int) Indicator {
	return rateOfChangeIndicator{
		series: series,
		window: window,
	}
}

func (roc rateOfChangeIndicator) Calculate(index int) big.Decimal {
	if index < roc.window {
		return big.ZERO
	}

	currentPrice := roc.series.Candles[index].ClosePrice
	oldPrice := roc.series.Candles[index-roc.window].ClosePrice

	// ROC = ((Current Price - Price n periods ago) / Price n periods ago) × 100
	return currentPrice.Sub(oldPrice).Div(oldPrice).Mul(big.NewDecimal(100))
}
