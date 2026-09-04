package techan

import "github.com/sdcoffey/big"

type commidityChannelIndexIndicator struct {
	series *TimeSeries
	window int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts *TimeSeries, window int) Indicator {
	requirePositiveWindow(window)
	return commidityChannelIndexIndicator{
		series: ts,
		window: window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) big.Decimal {
	if index < ccii.FirstValidIndex() {
		return big.ZERO
	}
	typicalPrice := NewTypicalPriceIndicator(ccii.series)
	typicalPriceSma := NewSimpleMovingAverage(typicalPrice, ccii.window)
	meanDeviation := NewMeanDeviationIndicator(typicalPrice, ccii.window).Calculate(index)
	if meanDeviation.IsZero() {
		return big.ZERO
	}

	return typicalPrice.Calculate(index).Sub(typicalPriceSma.Calculate(index)).Div(meanDeviation.Mul(big.NewFromString("0.015")))
}
