package techan

import "testing"

func TestAroonUpIndicator(t *testing.T) {
	t.Run("with < window periods", func(t *testing.T) {
		series := NewTimeSeries()
		indicator := NewHighPriceIndicator(series)

		aroonUp := NewAroonUpIndicator(indicator, 10)
		decimalEquals(t, 0, aroonUp.Calculate(0))
	})

	t.Run("with > window periods", func(t *testing.T) {
		ts := mockTimeSeriesFl(1, 2, 3, 4, 3, 2, 1)
		indicator := NewHighPriceIndicator(ts)

		aroonUpIndicator := NewAroonUpIndicator(indicator, 4)

		decimalEquals(t, 100, aroonUpIndicator.Calculate(3))
		decimalEquals(t, 75, aroonUpIndicator.Calculate(4))
		decimalEquals(t, 50, aroonUpIndicator.Calculate(5))
	})
}

func TestAroonDownIndicator(t *testing.T) {
	t.Run("with < window periods", func(t *testing.T) {
		series := NewTimeSeries()
		indicator := NewHighPriceIndicator(series)

		aroonUp := NewAroonDownIndicator(indicator, 10)
		decimalEquals(t, 0, aroonUp.Calculate(0))
	})

	t.Run("with > window periods", func(t *testing.T) {
		ts := mockTimeSeriesFl(5, 4, 3, 2, 3, 4, 5)
		indicator := NewLowPriceIndicator(ts)

		aroonUpIndicator := NewAroonDownIndicator(indicator, 4)

		decimalEquals(t, 100, aroonUpIndicator.Calculate(3))
		decimalEquals(t, 75, aroonUpIndicator.Calculate(4))
		decimalEquals(t, 50, aroonUpIndicator.Calculate(5))
	})
}
