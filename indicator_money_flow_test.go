package techan

import (
	"math"
	"testing"
)

var series = mockTimeSeriesOCHLV(
	[]float64{10, 12, 12, 8, 1000},
	[]float64{11, 14, 14, 9, 1500},
	[]float64{10, 20, 24, 10, 1200},
	[]float64{9, 10, 11, 9, 1800},
	[]float64{11, 14, 14, 9, 2000},
	[]float64{9, 10, 11, 9, 1300},
)

func TestMoneyFlowIndexIndicator(t *testing.T) {
	indicator := NewMoneyFlowIndexIndicator(series, 3)

	expectedValues := []float64{
		0,
		0,
		100,
		69.0189,
		71.9917,
		44.3114,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestMoneyFlowRatioIndicator(t *testing.T) {
	indicator := NewMoneyFlowRatioIndicator(series, 3)

	expectedValues := []float64{
		0,
		0,
		math.Inf(1),
		2.2278,
		2.5704,
		0.7957,
	}

	indicatorEquals(t, expectedValues, indicator)
}
