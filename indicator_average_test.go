package techan

import (
	"fmt"
	"testing"
)

func TestAverageGainLossInputWarmup(t *testing.T) {
	for _, tc := range []struct {
		name   string
		make   func(Indicator, int) Indicator
		prices []float64
	}{
		{"gains", NewAverageGainsIndicator, []float64{10, 12, 14, 18, 14, 24, 8}},
		{"losses", NewAverageLossesIndicator, []float64{30, 28, 26, 22, 26, 16, 32}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// The EMA first becomes usable at index 2. The selected changes
			// thereafter are 3, 0, 4.75, 0 for both gains and mirrored losses.
			for _, window := range []struct {
				size int
				want []float64
			}{
				{1, []float64{0, 0, 0, 3, 0, 4.75, 0}},
				{3, []float64{0, 0, 0, 1.5, 1, 31.0 / 12, 19.0 / 12}},
				{6, []float64{0, 0, 0, 1.5, 1, 31.0 / 16, 31.0 / 20}},
			} {
				t.Run(fmt.Sprint(window.size), func(t *testing.T) {
					average := tc.make(NewEMAIndicator(NewFixedIndicator(tc.prices...), 3), window.size)
					for index, want := range window.want {
						t.Run(fmt.Sprint(index), func(t *testing.T) {
							decimalEquals(t, want, average.Calculate(index))
						})
					}
				})
			}
		})
	}
}

func TestAverageGainsIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := mockTimeSeriesFl(1, 2, 3, 5, 8, 13)

		avgGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, avgGains.Calculate(0))
		decimalEquals(t, 1.0/2.0, avgGains.Calculate(1))
		decimalEquals(t, 2.0/3.0, avgGains.Calculate(2))
		decimalEquals(t, 1.0, avgGains.Calculate(3))
		decimalEquals(t, 7.0/5.0, avgGains.Calculate(4))
		decimalEquals(t, 12.0/6.0, avgGains.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := mockTimeSeriesFl(0, 5, 2, 10, 12, 11)

		cumGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumGains.Calculate(0))
		decimalEquals(t, 5/2.0, cumGains.Calculate(1))
		decimalEquals(t, 5/3.0, cumGains.Calculate(2))
		decimalEquals(t, 13.0/4.0, cumGains.Calculate(3))
		decimalEquals(t, 15.0/5.0, cumGains.Calculate(4))
		decimalEquals(t, 15.0/6.0, cumGains.Calculate(5))
	})

	t.Run("Rolling window", func(t *testing.T) {
		ts := mockTimeSeriesFl(0, 5, 2, 10, 12, 11)

		cumGains := NewAverageGainsIndicator(NewClosePriceIndicator(ts), 3)

		decimalEquals(t, 0, cumGains.Calculate(0))
		decimalEquals(t, 5.0/2.0, cumGains.Calculate(1))
		decimalEquals(t, 5.0/3.0, cumGains.Calculate(2))
		decimalEquals(t, 13.0/3.0, cumGains.Calculate(3))
		decimalEquals(t, 10.0/3.0, cumGains.Calculate(4))
		decimalEquals(t, 10.0/3.0, cumGains.Calculate(5))
	})
}

func TestNewAverageLossesIndicator(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 8, 5, 3, 2, 1)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 5.0/2.0, cumLosses.Calculate(1))
		decimalEquals(t, 8.0/3.0, cumLosses.Calculate(2))
		decimalEquals(t, 10.0/4.0, cumLosses.Calculate(3))
		decimalEquals(t, 11.0/5.0, cumLosses.Calculate(4))
		decimalEquals(t, 12.0/6.0, cumLosses.Calculate(5))
	})

	t.Run("Oscillating indicator", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 16, 10, 8, 9, 8)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 6)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 0, cumLosses.Calculate(1))
		decimalEquals(t, 6.0/3.0, cumLosses.Calculate(2))
		decimalEquals(t, 8.0/4.0, cumLosses.Calculate(3))
		decimalEquals(t, 8.0/5.0, cumLosses.Calculate(4))
		decimalEquals(t, 9.0/6.0, cumLosses.Calculate(5))
	})

	t.Run("Rolling window", func(t *testing.T) {
		ts := mockTimeSeriesFl(13, 16, 10, 8, 9, 8)

		cumLosses := NewAverageLossesIndicator(NewClosePriceIndicator(ts), 3)

		decimalEquals(t, 0, cumLosses.Calculate(0))
		decimalEquals(t, 0, cumLosses.Calculate(1))
		decimalEquals(t, 6.0/3.0, cumLosses.Calculate(2))
		decimalEquals(t, 8.0/3.0, cumLosses.Calculate(3))
		decimalEquals(t, 8.0/3.0, cumLosses.Calculate(4))
		decimalEquals(t, 1.0, cumLosses.Calculate(5))
	})
}
