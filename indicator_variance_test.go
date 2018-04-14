package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarianceIndicator(t *testing.T) {
	t.Run("when index is less than 1, returns 0", func(t *testing.T) {
		series := mockTimeSeries("0", "10")
		stdDev := StandardDeviationIndicator{
			Indicator: NewClosePriceIndicator(series),
		}

		assert.EqualValues(t, "0", stdDev.Calculate(0).String())
	})

	t.Run("returns the standard deviation when index > 2", func(t *testing.T) {
		series := mockTimeSeriesFl(
			10,
			2,
			38,
			23,
			38,
			23,
			21)

		stdDev := VarianceIndicator{
			Indicator: NewClosePriceIndicator(series),
		}

		assert.EqualValues(t, "16.00", stdDev.Calculate(1).FormattedString(2))
		assert.EqualValues(t, "238.22", stdDev.Calculate(2).FormattedString(2))
		assert.EqualValues(t, "186.19", stdDev.Calculate(3).FormattedString(2))
		assert.EqualValues(t, "211.36", stdDev.Calculate(4).FormattedString(2))
		assert.EqualValues(t, "176.22", stdDev.Calculate(5).FormattedString(2))
		assert.EqualValues(t, "151.27", stdDev.Calculate(6).FormattedString(2))
	})
}
