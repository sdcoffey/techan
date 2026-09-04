package techan

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	// Wilder seed: the first three actual price changes, followed by 1/3 smoothing.
	expectedValues := []float64{0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 53.6424, 50.2715, 20.8260, 43.8487, 14.0604, 21.1501}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 1.1571, 1.0109, 0.2630, 0.7809, 0.1636, 0.2682}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.ZERO.FormattedString(2), rsInd.Calculate(2).FormattedString(2))
}
