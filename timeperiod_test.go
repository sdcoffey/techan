package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	t.Run("SimpleDT:SimpleDT", func(t *testing.T) {
		parseable := "01/20/2009T12:00:00:01/20/2017T12:00:00"
		TimePeriod, err := Parse(parseable)
		assert.NoError(t, err)

		assert.EqualValues(t, 2009, TimePeriod.Start.Year())
		assert.EqualValues(t, 01, TimePeriod.Start.Month())
		assert.EqualValues(t, 20, TimePeriod.Start.Day())
		assert.EqualValues(t, 12, TimePeriod.Start.Hour())
		assert.EqualValues(t, 0, TimePeriod.Start.Minute())
		assert.EqualValues(t, 0, TimePeriod.Start.Second())

		assert.EqualValues(t, 2017, TimePeriod.End.Year())
		assert.EqualValues(t, 01, TimePeriod.End.Month())
		assert.EqualValues(t, 20, TimePeriod.End.Day())
		assert.EqualValues(t, 12, TimePeriod.End.Hour())
		assert.EqualValues(t, 0, TimePeriod.End.Minute())
		assert.EqualValues(t, 0, TimePeriod.End.Second())
	})

	// this has the potential to be flaky, on account of the slight difference in time between
	// now and the now created in Parse
	t.Run("SimpleDT:", func(t *testing.T) {
		parseable := "08/15/1991T20:30:00:"
		TimePeriod, err := Parse(parseable)
		now := time.Now()
		assert.NoError(t, err)

		assert.EqualValues(t, 1991, TimePeriod.Start.Year())
		assert.EqualValues(t, 8, TimePeriod.Start.Month())
		assert.EqualValues(t, 15, TimePeriod.Start.Day())
		assert.EqualValues(t, 20, TimePeriod.Start.Hour())
		assert.EqualValues(t, 30, TimePeriod.Start.Minute())
		assert.EqualValues(t, 0, TimePeriod.Start.Second())

		assert.EqualValues(t, now.Year(), TimePeriod.End.Year())
		assert.EqualValues(t, now.Month(), TimePeriod.End.Month())
		assert.EqualValues(t, now.Day(), TimePeriod.End.Day())
		assert.EqualValues(t, now.Hour(), TimePeriod.End.Hour())
		assert.EqualValues(t, now.Minute(), TimePeriod.End.Minute())
		assert.EqualValues(t, now.Second(), TimePeriod.End.Second())
	})

	t.Run("SimpleDate:SimpleDate", func(t *testing.T) {
		parseable := "09/01/1773:07/04/1776"
		TimePeriod, err := Parse(parseable)
		assert.NoError(t, err)

		assert.EqualValues(t, 1773, TimePeriod.Start.Year())
		assert.EqualValues(t, 9, TimePeriod.Start.Month())
		assert.EqualValues(t, 1, TimePeriod.Start.Day())

		assert.EqualValues(t, 1776, TimePeriod.End.Year())
		assert.EqualValues(t, 7, TimePeriod.End.Month())
		assert.EqualValues(t, 4, TimePeriod.End.Day())
	})

	t.Run("SimpleDate:", func(t *testing.T) {
		parseable := "07/04/1776:"
		TimePeriod, err := Parse(parseable)
		now := time.Now()
		assert.NoError(t, err)

		assert.EqualValues(t, 1776, TimePeriod.Start.Year())
		assert.EqualValues(t, 7, TimePeriod.Start.Month())
		assert.EqualValues(t, 4, TimePeriod.Start.Day())

		assert.EqualValues(t, now.Year(), TimePeriod.End.Year())
		assert.EqualValues(t, now.Month(), TimePeriod.End.Month())
		assert.EqualValues(t, now.Day(), TimePeriod.End.Day())
	})
}

func TestTimePeriod_Length(t *testing.T) {
	now := time.Now()
	TimePeriod := NewTimePeriod(now.Add(-time.Minute*10), now)

	assert.EqualValues(t, time.Minute*10, TimePeriod.Length())
}

func TestTimePeriod_Since(t *testing.T) {
	now := time.Now()

	t.Run("0", func(t *testing.T) {
		tp := NewTimePeriod(now, now.Add(time.Minute))
		previousTimePeriod := NewTimePeriod(now.Add(-time.Minute), now)

		since := tp.Since(previousTimePeriod)
		assert.EqualValues(t, 0, since)
	})

	t.Run("Positive", func(t *testing.T) {
		tp := NewTimePeriod(now, now.Add(time.Minute))
		previousTimePeriod := NewTimePeriod(now.Add(-time.Minute*2), now.Add(-time.Minute))

		since := tp.Since(previousTimePeriod)

		assert.EqualValues(t, time.Minute, since)
	})
}
