package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCumulativeGainsIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumGains.Calculate(0))
		assert.EqualValues(t, 1, cumGains.Calculate(1))
		assert.EqualValues(t, 2, cumGains.Calculate(2))
		assert.EqualValues(t, 4, cumGains.Calculate(3))
		assert.EqualValues(t, 7, cumGains.Calculate(4))
		assert.EqualValues(t, 12, cumGains.Calculate(5))
	})

	t.Run("Oscillating scale", func(t *testing.T) {
		ts := MockTimeSeries(0, 5, 2, 10, 12, 11)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumGains.Calculate(0))
		assert.EqualValues(t, 5, cumGains.Calculate(1))
		assert.EqualValues(t, 5, cumGains.Calculate(2))
		assert.EqualValues(t, 13, cumGains.Calculate(3))
		assert.EqualValues(t, 15, cumGains.Calculate(4))
		assert.EqualValues(t, 15, cumGains.Calculate(5))
	})

	t.Run("Rolling timeframe", func(t *testing.T) {
		ts := MockTimeSeries(0, 5, 2, 10, 12, 11)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 3)

		assert.EqualValues(t, 0, cumGains.Calculate(0))
		assert.EqualValues(t, 5, cumGains.Calculate(1))
		assert.EqualValues(t, 5, cumGains.Calculate(2))
		assert.EqualValues(t, 13, cumGains.Calculate(3))
		assert.EqualValues(t, 10, cumGains.Calculate(4))
		assert.EqualValues(t, 10, cumGains.Calculate(5))
	})
}

func TestCumulativeLossesIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := MockTimeSeries(13, 8, 5, 3, 2, 1)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 5, cumLosses.Calculate(1))
		assert.EqualValues(t, 8, cumLosses.Calculate(2))
		assert.EqualValues(t, 10, cumLosses.Calculate(3))
		assert.EqualValues(t, 11, cumLosses.Calculate(4))
		assert.EqualValues(t, 12, cumLosses.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := MockTimeSeries(13, 16, 10, 8, 9, 8)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 0, cumLosses.Calculate(1))
		assert.EqualValues(t, 6, cumLosses.Calculate(2))
		assert.EqualValues(t, 8, cumLosses.Calculate(3))
		assert.EqualValues(t, 8, cumLosses.Calculate(4))
		assert.EqualValues(t, 9, cumLosses.Calculate(5))
	})

	t.Run("Rolling timeframe", func(t *testing.T) {
		ts := MockTimeSeries(13, 16, 10, 8, 9, 8)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 3)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 0, cumLosses.Calculate(1))
		assert.EqualValues(t, 6, cumLosses.Calculate(2))
		assert.EqualValues(t, 8, cumLosses.Calculate(3))
		assert.EqualValues(t, 8, cumLosses.Calculate(4))
		assert.EqualValues(t, 3, cumLosses.Calculate(5))
	})
}

func TestPercentGainIndicator(t *testing.T) {
	t.Run("Up", func(t *testing.T) {
		ts := MockTimeSeries(1, 1.5, 2.25, 2.25)

		pgi := NewPercentChangeIndicator(NewClosePriceIndicator(ts))

		assert.EqualValues(t, 0, pgi.Calculate(0))
		assert.EqualValues(t, .5, pgi.Calculate(1))
		assert.EqualValues(t, .5, pgi.Calculate(2))
		assert.EqualValues(t, 0, pgi.Calculate(3))
	})

	t.Run("Down", func(t *testing.T) {
		ts := MockTimeSeries(10, 5, 2.5, 2.5)

		pgi := NewPercentChangeIndicator(NewClosePriceIndicator(ts))

		assert.EqualValues(t, 0, pgi.Calculate(0))
		assert.EqualValues(t, -.5, pgi.Calculate(1))
		assert.EqualValues(t, -.5, pgi.Calculate(2))
		assert.EqualValues(t, 0, pgi.Calculate(3))
	})
}
