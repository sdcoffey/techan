package techan

import (
	"testing"
)

func TestGainIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(1, 2, 3, 3, 2, 1)

	gains := NewGainIndicator(NewClosePriceIndicator(ts))

	decimalEquals(t, 0, gains.Calculate(0))
	decimalEquals(t, 1, gains.Calculate(1))
	decimalEquals(t, 1, gains.Calculate(2))
	decimalEquals(t, 0, gains.Calculate(3))
	decimalEquals(t, 0, gains.Calculate(4))
	decimalEquals(t, 0, gains.Calculate(5))
}

func TestLossIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(1, 2, 3, 3, 2, 0)

	gains := NewLossIndicator(NewClosePriceIndicator(ts))

	decimalEquals(t, 0, gains.Calculate(0))
	decimalEquals(t, 0, gains.Calculate(1))
	decimalEquals(t, 0, gains.Calculate(2))
	decimalEquals(t, 0, gains.Calculate(3))
	decimalEquals(t, 1, gains.Calculate(4))
	decimalEquals(t, 2, gains.Calculate(5))
}

func TestCumulativeGainsIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := mockTimeSeriesFl(1, 2, 3, 5, 8, 13)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumGains.Calculate(0))
		decimalEquals(t, 1, cumGains.Calculate(1))
		decimalEquals(t, 2, cumGains.Calculate(2))
		decimalEquals(t, 4, cumGains.Calculate(3))
		decimalEquals(t, 7, cumGains.Calculate(4))
		decimalEquals(t, 12, cumGains.Calculate(5))
	})

	t.Run("Oscillating scale", func(t *testing.T) {
		ts := mockTimeSeriesFl(0, 5, 2, 10, 12, 11)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumGains.Calculate(0))
		decimalEquals(t, 5, cumGains.Calculate(1))
		decimalEquals(t, 5, cumGains.Calculate(2))
		decimalEquals(t, 13, cumGains.Calculate(3))
		decimalEquals(t, 15, cumGains.Calculate(4))
		decimalEquals(t, 15, cumGains.Calculate(5))
	})

	t.Run("Rolling timeframe", func(t *testing.T) {
		ts := mockTimeSeriesFl(0, 5, 2, 10, 12, 11)

		cumGains := NewCumulativeGainsIndicator(NewClosePriceIndicator(ts), 3)

		decimalEquals(t, 0, cumGains.Calculate(0))
		decimalEquals(t, 5, cumGains.Calculate(1))
		decimalEquals(t, 5, cumGains.Calculate(2))
		decimalEquals(t, 13, cumGains.Calculate(3))
		decimalEquals(t, 10, cumGains.Calculate(4))
		decimalEquals(t, 10, cumGains.Calculate(5))
	})
}

func TestCumulativeLossesIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 8, 5, 3, 2, 1)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 5, cumLosses.Calculate(1))
		decimalEquals(t, 8, cumLosses.Calculate(2))
		decimalEquals(t, 10, cumLosses.Calculate(3))
		decimalEquals(t, 11, cumLosses.Calculate(4))
		decimalEquals(t, 12, cumLosses.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 16, 10, 8, 9, 8)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 0, cumLosses.Calculate(1))
		decimalEquals(t, 6, cumLosses.Calculate(2))
		decimalEquals(t, 8, cumLosses.Calculate(3))
		decimalEquals(t, 8, cumLosses.Calculate(4))
		decimalEquals(t, 9, cumLosses.Calculate(5))
	})

	t.Run("Rolling timeframe", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 16, 10, 8, 9, 8)

		cumLosses := NewCumulativeLossesIndicator(NewClosePriceIndicator(ts), 3)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 0, cumLosses.Calculate(1))
		decimalEquals(t, 6, cumLosses.Calculate(2))
		decimalEquals(t, 8, cumLosses.Calculate(3))
		decimalEquals(t, 8, cumLosses.Calculate(4))
		decimalEquals(t, 3, cumLosses.Calculate(5))
	})
}

func TestPercentGainIndicator(t *testing.T) {
	t.Run("Up", func(t *testing.T) {
		ts := mockTimeSeriesFl(1, 1.5, 2.25, 2.25)

		pgi := NewPercentChangeIndicator(NewClosePriceIndicator(ts))

		decimalEquals(t, 0, pgi.Calculate(0))
		decimalEquals(t, .5, pgi.Calculate(1))
		decimalEquals(t, .5, pgi.Calculate(2))
		decimalEquals(t, 0, pgi.Calculate(3))
	})

	t.Run("Down", func(t *testing.T) {
		ts := mockTimeSeriesFl(10, 5, 2.5, 2.5)

		pgi := NewPercentChangeIndicator(NewClosePriceIndicator(ts))

		decimalEquals(t, 0, pgi.Calculate(0))
		decimalEquals(t, -.5, pgi.Calculate(1))
		decimalEquals(t, -.5, pgi.Calculate(2))
		decimalEquals(t, 0, pgi.Calculate(3))
	})
}
