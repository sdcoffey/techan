package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardDeviationIndicator(t *testing.T) {
	t.Run("when index is less than 1, returns 0", func(t *testing.T) {
		series := mockTimeSeries("0", "10")
		stdDev := NewStandardDeviationIndicator(NewClosePriceIndicator(series))

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

		stdDev := NewStandardDeviationIndicator(NewClosePriceIndicator(series))

		assert.EqualValues(t, "4.00", stdDev.Calculate(1).FormattedString(2))
		assert.EqualValues(t, "15.43", stdDev.Calculate(2).FormattedString(2))
		assert.EqualValues(t, "13.65", stdDev.Calculate(3).FormattedString(2))
		assert.EqualValues(t, "14.54", stdDev.Calculate(4).FormattedString(2))
		assert.EqualValues(t, "13.27", stdDev.Calculate(5).FormattedString(2))
		assert.EqualValues(t, "12.30", stdDev.Calculate(6).FormattedString(2))
	})
}
