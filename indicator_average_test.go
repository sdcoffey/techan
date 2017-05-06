package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverageGainsIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := MockTimeSeries(1, 2, 3, 5, 8, 13)

		avgGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, avgGains.Calculate(0))
		assert.EqualValues(t, 1.0/2.0, avgGains.Calculate(1))
		assert.EqualValues(t, 2.0/3.0, avgGains.Calculate(2))
		assert.EqualValues(t, 4.0/4.0, avgGains.Calculate(3))
		assert.EqualValues(t, 7.0/5.0, avgGains.Calculate(4))
		assert.EqualValues(t, 12.0/6.0, avgGains.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := MockTimeSeries(0, 5, 2, 10, 12, 11)

		cumGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumGains.Calculate(0))
		assert.EqualValues(t, 5/2.0, cumGains.Calculate(1))
		assert.EqualValues(t, 5/3.0, cumGains.Calculate(2))
		assert.EqualValues(t, 13.0/4.0, cumGains.Calculate(3))
		assert.EqualValues(t, 15.0/5.0, cumGains.Calculate(4))
		assert.EqualValues(t, 15.0/6.0, cumGains.Calculate(5))
	})

	t.Run("Rolling window", func(t *testing.T) {
		ts := MockTimeSeries(0, 5, 2, 10, 12, 11)

		cumGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 3)

		assert.EqualValues(t, 0, cumGains.Calculate(0))
		assert.EqualValues(t, 5.0/2.0, cumGains.Calculate(1))
		assert.EqualValues(t, 5.0/3.0, cumGains.Calculate(2))
		assert.EqualValues(t, 13.0/3.0, cumGains.Calculate(3))
		assert.EqualValues(t, 10.0/3.0, cumGains.Calculate(4))
		assert.EqualValues(t, 10.0/3.0, cumGains.Calculate(5))
	})
}

func TestNewAverageLossesIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := MockTimeSeries(13, 8, 5, 3, 2, 1)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 5.0/2.0, cumLosses.Calculate(1))
		assert.EqualValues(t, 8.0/3.0, cumLosses.Calculate(2))
		assert.EqualValues(t, 10.0/4.0, cumLosses.Calculate(3))
		assert.EqualValues(t, 11.0/5.0, cumLosses.Calculate(4))
		assert.EqualValues(t, 12.0/6.0, cumLosses.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := MockTimeSeries(13, 16, 10, 8, 9, 8)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 6)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 0, cumLosses.Calculate(1))
		assert.EqualValues(t, 6.0/3.0, cumLosses.Calculate(2))
		assert.EqualValues(t, 8.0/4.0, cumLosses.Calculate(3))
		assert.EqualValues(t, 8.0/5.0, cumLosses.Calculate(4))
		assert.EqualValues(t, 9.0/6.0, cumLosses.Calculate(5))
	})

	t.Run("Rolling window", func(t *testing.T) {
		ts := MockTimeSeries(13, 16, 10, 8, 9, 8)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 3)

		assert.EqualValues(t, 0, cumLosses.Calculate(0))
		assert.EqualValues(t, 0, cumLosses.Calculate(1))
		assert.EqualValues(t, 6.0/3.0, cumLosses.Calculate(2))
		assert.EqualValues(t, 8.0/3.0, cumLosses.Calculate(3))
		assert.EqualValues(t, 8.0/3.0, cumLosses.Calculate(4))
		assert.EqualValues(t, 3.0/3.0, cumLosses.Calculate(5))
	})
}
