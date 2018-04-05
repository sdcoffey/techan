package talib4g

import (
	"fmt"
	"time"
)

// TimePeriod is a simple struct that describes a period of time with a Start and End time
type TimePeriod struct {
	Start time.Time
	End   time.Time
}

// Constants representing basic, human-readable and writable date formats
const (
	SimpleDateTimeFormat = "01/02/2006T15:04:05"
	SimpleDateFormat     = "01/02/2006"
)

// Parse takes a string in one of the following formats and returns a new TimePeriod, and optionally, an error
//
// Supported Formats:
// SimpleDateTimeFormat:SimpleDateTimeFormat
// SimpleDateTimeFormat: (to now)
// SimpleDateFormat:
// SimpleDateFormat:SimpleDateFormat
func Parse(timerange string) (tr TimePeriod, err error) {
	var start, end, layout string
	switch len(timerange) {
	case len(SimpleDateTimeFormat)*2 + 1:
		layout = SimpleDateTimeFormat
		start = string(timerange[:len(SimpleDateTimeFormat)])
		end = string(timerange[len(SimpleDateTimeFormat)+1:])
	case len(SimpleDateTimeFormat) + 1:
		layout = SimpleDateTimeFormat
		start = string(timerange[:len(SimpleDateTimeFormat)])
		end = ""
	case len(SimpleDateFormat)*2 + 1:
		layout = SimpleDateFormat
		start = string(timerange[:len(SimpleDateFormat)])
		end = string(timerange[len(SimpleDateFormat)+1:])
	case len(SimpleDateFormat) + 1:
		layout = SimpleDateFormat
		start = string(timerange[:len(SimpleDateFormat)])
		end = ""
	default:
		err = fmt.Errorf("could not parse timerange string %s", timerange)
		return
	}

	if tr.Start, err = time.Parse(layout, start); err != nil {
		err = fmt.Errorf("could not parse time string %s", start)
	}

	if end == "" {
		tr.End = time.Now()
	} else if tr.End, err = time.Parse(layout, end); err != nil {
		err = fmt.Errorf("could not parse time string %s", end)
	}

	return
}

// Length returns the length of the period as a time.Duration value
func (tp TimePeriod) Length() time.Duration {
	return tp.End.Sub(tp.Start)
}

// Since returns the amount of time elapsed since the end of another TimePeriod as a time.Duration value
func (tp TimePeriod) Since(other TimePeriod) time.Duration {
	return tp.Start.Sub(other.End)
}

// Format returns the string representaion of this timePeriod in the given format
func (tp TimePeriod) Format(layout string) string {
	return fmt.Sprintf("%s -> %s", tp.Start.Format(layout), tp.End.Format(layout))
}

// Advance will return a new TimePeriod with the start and end periods moved forwards or backwards in time in accordance
// with the number of iterations given.
//
// Example:
// A timePeriod that is one hour long, starting at unix time 0 and ending at unix time 3600, and advanced by one,
// will return a time period starting at unix time 3600 and ending at unix time 7200
func (tp TimePeriod) Advance(iterations int) TimePeriod {
	return TimePeriod{
		Start: tp.Start.Add(tp.Length() * time.Duration(iterations)),
		End:   tp.End.Add(tp.Length() * time.Duration(iterations)),
	}
}

func (tp TimePeriod) String() string {
	return tp.Format(SimpleDateTimeFormat)
}

// NewTimePeriod returns a TimePeriod starting at the given time and ending at the given time plus the given duration
func NewTimePeriod(start time.Time, period time.Duration) TimePeriod {
	return TimePeriod{
		Start: start,
		End:   start.Add(period),
	}
}
