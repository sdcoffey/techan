package techan

import "testing"

func TestMaximumDrawdownIndicator(t *testing.T) {
	t.Run("with window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 0, 20, 1, 4)

		mvi := NewMaximumDrawdownIndicator(NewClosePriceIndicator(ts), 3)
		decimalEquals(t, -0.95, mvi.Calculate(ts.LastIndex()))
	})

	t.Run("without window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 0, 20, 1, 4)

		mvi := NewMaximumDrawdownIndicator(NewClosePriceIndicator(ts), -1)
		decimalEquals(t, -1.05, mvi.Calculate(ts.LastIndex()))
	})
}
