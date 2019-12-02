package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var rsTestMockSeries = mockTimeSeries(
	"44.34",
	"44.09",
	"44.15",
	"43.61",
	"44.33",
	"44.83",
	"45.10",
	"45.42",
	"45.84",
	"46.08",
	"45.89",
	"46.03",
	"45.61",
	"46.28",
	"46.28",
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	t.Run("when index == 0, returns 0", func(t *testing.T) {
		closeIndicator := NewClosePriceIndicator(rsTestMockSeries)
		indicator := NewRelativeStrengthIndexIndicator(closeIndicator, 14)

		assert.EqualValues(t, "0.00", indicator.Calculate(0).FormattedString(2))
	})

	t.Run("when index > 0, returns rsi index", func(t *testing.T) {
		closeIndicator := NewClosePriceIndicator(rsTestMockSeries)
		indicator := NewRelativeStrengthIndexIndicator(closeIndicator, 14)

		assert.EqualValues(t, "70.46", indicator.Calculate(14).FormattedString(2))
	})
}

func TestRelativeStrengthIndicator(t *testing.T) {
	closeIndicator := NewClosePriceIndicator(rsTestMockSeries)
	indicator := NewRelativeStrengthIndicator(closeIndicator, 14)

	assert.EqualValues(t, "2.39", indicator.Calculate(14).FormattedString(2))
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.MaxFloat64).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
