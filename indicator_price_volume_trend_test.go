package techan

import "testing"

func TestPrceVolumeTrendIndicator(t *testing.T) {
	indicator := NewPriceVolumeTrendIndicator(NewClosePriceIndicator(mockedTimeSeries),
		NewVolumeIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		-0.1795,
		-0.5375,
		0.1908,
		0.1308,
		-0.7565,
		-0.3337,
		-2.308,
		-2.1275,
	}

	indicatorEquals(t, expectedValues, indicator)
}
