package techan

import "testing"

func TestKeltnerChannel(t *testing.T) {
	t.Run("Upper", func(t *testing.T) {
		ts := mockTimeSeriesOCHL(
			[]float64{10, 15, 20, 10},
			[]float64{11, 16, 21, 11},
			[]float64{12, 17, 22, 12},
			[]float64{13, 18, 23, 13},
			[]float64{14, 19, 24, 14},
			[]float64{15, 20, 25, 15},
			[]float64{16, 20, 26, 16},
		)

		upper := NewKeltnerChannelUpperIndicator(ts, 3)

		decimalEquals(t, 40, upper.Calculate(5))
		decimalEquals(t, 39, upper.Calculate(4))
		decimalEquals(t, 38, upper.Calculate(3))
		decimalEquals(t, 37, upper.Calculate(2))
		decimalEquals(t, 0, upper.Calculate(1))
		decimalEquals(t, 0, upper.Calculate(0))
	})
}
