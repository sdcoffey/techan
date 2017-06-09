package talib4g

import (
	"fmt"
	"time"
)

type TimePeriod struct {
	Start time.Time
	End   time.Time
}

const SimpleDateTimeFormat = "01/02/2006T15:04:05"
const SimpleDateFormat = "01/02/2006"

// Support
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
		err = fmt.Errorf("Could not parse timerange string %s", timerange)
		return
	}

	if tr.Start, err = time.Parse(layout, start); err != nil {
		err = fmt.Errorf("Could not parse timerange string %s -> %s", timerange, err)
	}

	if end == "" {
		tr.End = time.Now()
	} else if tr.End, err = time.Parse(layout, end); err != nil {
		err = fmt.Errorf("Could not parse timerange string %s -> %s", timerange, err)
	}

	return
}

func (tp TimePeriod) Length() time.Duration {
	return tp.End.Sub(tp.Start)
}

func (tp TimePeriod) Since(other TimePeriod) time.Duration {
	return tp.Start.Sub(other.End)
}

func (tp TimePeriod) Format(layout string) string {
	return fmt.Sprintf("%s -> %s", tp.Start.Format(layout), tp.End.Format(layout))
}

func (tp TimePeriod) String() string {
	return tp.Format(SimpleDateTimeFormat)
}

func NewTimePeriod(start, end time.Time) TimePeriod {
	return TimePeriod{
		Start: start,
		End:   end,
	}
}

func NewTimePeriodD(start time.Time, period time.Duration) TimePeriod {
	return TimePeriod{
		Start: start,
		End:   start.Add(period),
	}
}
