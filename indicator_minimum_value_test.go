package techan

import "testing"

func TestMinimumValueIndicator(t *testing.T) {
	t.Run("with window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 0, 20, 1, 4)

		mvi := NewMinimumValueIndicator(NewClosePriceIndicator(ts), 3)
		decimalEquals(t, 1, mvi.Calculate(ts.LastIndex()))
		decimalEquals(t, 0, mvi.Calculate(ts.LastIndex()-1))
	})

	t.Run("without window", func(t *testing.T) {
		ts := mockTimeSeriesFl(-1, 10, 0, 20, 1, 4)

		mvi := NewMinimumValueIndicator(NewClosePriceIndicator(ts), -1)
		decimalEquals(t, -1, mvi.Calculate(ts.LastIndex()))
	})
}
