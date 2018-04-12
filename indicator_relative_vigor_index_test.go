package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeVigorIndexIndicator_Calculate(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]string{"10", "12", "12", "8"},
		[]string{"11", "14", "14", "9"},
		[]string{"8", "19", "20", "8"},
		[]string{"9", "10", "11", "8"},
	)

	rvii := NewRelativeVigorIndexIndicator(series)

	t.Run("Returns zero when index < 4", func(t *testing.T) {
		assert.EqualValues(t, "0", rvii.Calculate(0).String())
		assert.EqualValues(t, "0", rvii.Calculate(1).String())
		assert.EqualValues(t, "0", rvii.Calculate(2).String())
	})

	t.Run("Calculates rvii", func(t *testing.T) {
		assert.EqualValues(t, "0.756", rvii.Calculate(3).FormattedString(3))
	})
}

func TestRelativeVigorIndexSignalLine_Calculate(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]string{"10", "12", "12", "8"},
		[]string{"11", "14", "14", "9"},
		[]string{"8", "19", "20", "8"},
		[]string{"9", "10", "11", "8"},
		[]string{"11", "14", "14", "9"},
		[]string{"9", "10", "11", "8"},
		[]string{"10", "12", "12", "8"},
		[]string{"9", "10", "11", "8"},
	)

	signalLine := NewRelativeVigorSignalLine(series)

	t.Run("Returns zero when index < 0", func(t *testing.T) {
		assert.EqualValues(t, "0", signalLine.Calculate(0).String())
		assert.EqualValues(t, "0", signalLine.Calculate(1).String())
		assert.EqualValues(t, "0", signalLine.Calculate(2).String())
	})

	t.Run("Calculates rvii signal line", func(t *testing.T) {
		assert.EqualValues(t, "0.5752", signalLine.Calculate(7).FormattedString(4))
	})
}
