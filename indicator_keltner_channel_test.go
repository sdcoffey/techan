package techan

import (
	"testing"
)

func TestKeltnerChannel(t *testing.T) {
	t.Run("Upper", func(t *testing.T) {
		upper := NewKeltnerChannelUpperIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			67.91,
			67.73,
			67.46,
			67.685,
			67.7675,
			67.3588,
			67.3644,
			67.0405,
			66.6219,
		}

		indicatorEquals(t, expectedValues, upper)
	})

	t.Run("Lower", func(t *testing.T) {
		lower := NewKeltnerChannelLowerIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			59.91,
			59.73,
			59.46,
			59.685,
			59.7675,
			59.3588,
			59.3644,
			57.6539,
			57.2353,
		}

		indicatorEquals(t, expectedValues, lower)
	})
}
