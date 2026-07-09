package techan

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestStochasticRSIIndicator(t *testing.T) {
	indicator := NewStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries), 5)

	expectedValues := []float64{
		100,
		100,
		100,
		100,
		100,
		100,
		100,
		95.9481,
		54.5245,
		93.1791,
		0,
		21.6754,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestFastStochasticRSIIndicator(t *testing.T) {
	indicator := NewFastStochasticRSIIndicator(NewStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries),
		5), 3)

	expectedValues := []float64{
		0,
		0,
		100,
		100,
		100,
		100,
		100,
		98.6494,
		83.4909,
		81.2173,
		49.2346,
		38.2848,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestSlowStochasticRSIIndicator(t *testing.T) {
	indicator := NewSlowStochasticRSIIndicator(NewFastStochasticRSIIndicator(NewStochasticRSIIndicator(
		NewClosePriceIndicator(mockedTimeSeries), 5), 3), 3)

	expectedValues := []float64{
		0,
		0,
		33.3333,
		66.6667,
		100,
		100,
		100,
		99.5498,
		94.0468,
		87.7858,
		71.3142,
		56.2456,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestFastStochasticRSIIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewStochasticRSIIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(100).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
