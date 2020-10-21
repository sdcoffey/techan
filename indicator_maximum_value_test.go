package techan

import "testing"

func TestMaximumValueIndicator(t *testing.T) {
	t.Run("with window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 21, 20, 1, 4)

		mvi := NewMaximumValueIndicator(NewClosePriceIndicator(ts), 3)
		decimalEquals(t, 20, mvi.Calculate(ts.LastIndex()))
		decimalEquals(t, 21, mvi.Calculate(ts.LastIndex()-1))
	})

	t.Run("without window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 0, 20, 1, 4)

		mvi := NewMaximumValueIndicator(NewClosePriceIndicator(ts), -1)
		decimalEquals(t, 20, mvi.Calculate(ts.LastIndex()))
	})
}
