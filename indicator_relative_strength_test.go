package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var timeseries = mockTimeSeriesFl(
	50.45, 50.30, 50.20, 50.15,
	50.05, 50.06, 50.10, 50.08,
	50.03, 50.07, 50.01, 50.14,
	50.22, 50.43, 50.50, 50.56,
	50.52, 50.70, 50.55, 50.62,
	50.90, 50.82, 50.86, 51.20,
	51.30, 51.10)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	t.Run("when index == 0, returns 0", func(t *testing.T) {
		indicator := NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(timeseries), 14)

		decimalEquals(t, 0, indicator.Calculate(0))
	})

	t.Run("when index > 0, returns rsi index", func(t *testing.T) {
		indicator := NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(timeseries), 14)

		decimalEquals(t, 68.4747, indicator.Calculate(15))
		decimalEquals(t, 64.7836, indicator.Calculate(16))
		decimalEquals(t, 72.0777, indicator.Calculate(17))
		decimalEquals(t, 60.7800, indicator.Calculate(18))
		decimalEquals(t, 63.6439, indicator.Calculate(19))
		decimalEquals(t, 72.3434, indicator.Calculate(20))
		decimalEquals(t, 67.3823, indicator.Calculate(21))
		decimalEquals(t, 68.5438, indicator.Calculate(22))
		decimalEquals(t, 76.2770, indicator.Calculate(23))
		decimalEquals(t, 77.9908, indicator.Calculate(24))
		decimalEquals(t, 67.4895, indicator.Calculate(25))
	})
}

func TestRelativeStrengthIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndicator(NewClosePriceIndicator(timeseries), 14)

	decimalEquals(t, 2.1721, indicator.Calculate(15))
	decimalEquals(t, 1.8396, indicator.Calculate(16))
	decimalEquals(t, 2.5814, indicator.Calculate(17))
	decimalEquals(t, 1.5497, indicator.Calculate(18))
	decimalEquals(t, 1.7506, indicator.Calculate(19))
	decimalEquals(t, 2.6158, indicator.Calculate(20))
	decimalEquals(t, 2.0658, indicator.Calculate(21))
	decimalEquals(t, 2.1790, indicator.Calculate(22))
	decimalEquals(t, 3.2153, indicator.Calculate(23))
	decimalEquals(t, 3.5436, indicator.Calculate(24))
	decimalEquals(t, 2.0759, indicator.Calculate(25))
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.MaxFloat64).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
