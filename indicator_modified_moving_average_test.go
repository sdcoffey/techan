package techan

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestModifiedMovingAverage(t *testing.T) {
	indicator := NewMMAIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expected := []float64{
		0,
		0,
		64.09,
		63.97,
		63.83,
		63.6167,
		63.7144,
		63.7596,
		63.4898,
		63.4498,
		62.7432,
		62.3321,
	}

	indicatorEquals(t, expected, indicator)
}

func TestModifiedMovingAverage_ResetCacheFrom(t *testing.T) {
	series := mockTimeSeriesFl(10, 10, 10, 10)
	mma := NewMMAIndicator(NewClosePriceIndicator(series), 3)

	decimalEquals(t, 10, mma.Calculate(3))

	series.Candles[3].ClosePrice = big.NewFromString("20")
	decimalEquals(t, 10, mma.Calculate(3))

	assert.True(t, ResetCacheFrom(mma, 3))
	decimalEquals(t, 13.3333, mma.Calculate(3))
}
