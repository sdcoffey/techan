package techan

// NewMACDIndicator returns a derivative Indicator which returns the difference between two EMAIndicators with long and
// short windows. It's useful for gauging the strength of price movements. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/m/macd.asp
func NewMACDIndicator(baseIndicator Indicator, shortwindow, longwindow int) Indicator {
	return NewDifferenceIndicator(NewEMAIndicator(baseIndicator, shortwindow), NewEMAIndicator(baseIndicator, longwindow))
}

// NewMACDHistogramIndicator returns a derivative Indicator based on the MACDIndicator, the result of which is
// the macd indicator minus it's signalLinewindow EMA. A more in-depth explanation can be found here:
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:macd-histogram
func NewMACDHistogramIndicator(macdIdicator Indicator, signalLinewindow int) Indicator {
	return NewDifferenceIndicator(macdIdicator, NewEMAIndicator(macdIdicator, signalLinewindow))
}
