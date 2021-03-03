package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMACDIndicator(t *testing.T) {
	series := randomTimeSeries(100)

	macd := NewMACDIndicator(NewClosePriceIndicator(series), 12, 26)

	assert.NotNil(t, macd)
}

func TestNewMACDHistogramIndicator(t *testing.T) {
	series := randomTimeSeries(100)

	macd := NewMACDIndicator(NewClosePriceIndicator(series), 12, 26)
	macdHistogram := NewMACDHistogramIndicator(macd, 9)

	assert.NotNil(t, macdHistogram)
}
