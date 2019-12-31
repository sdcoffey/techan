package techan

import (
	"fmt"
	"regexp"
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

	SimpleTimeFormat   = "15:04:05"
	SimpleDateFormatV2 = "2006-01-02"
)

// Constants representing regexes for parsing datetimes
var (
	SimpleTimeFomatRegex    = regexp.MustCompile(`T\d{2}:\d{2}:\d{2}`)
	SimpleDateFormatV2Regex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

// ParseTimePeriod parses two datetimes as one string and returns it as a TimePeriod.
//
// Note that if you were previously using Parse, the date format has changed to something more rfc3339-like (yyyy-mm-dd)
// Will accept any combination of date and time for either side. Omitting the right hand side will result in a time
// period ending in time.Now()
func ParseTimePeriod(period string) (TimePeriod, error) {
	dateMatches := SimpleDateFormatV2Regex.FindAllStringIndex(period, -1)
	timeMatches := SimpleTimeFomatRegex.FindAllStringIndex(period, -1)

	formats := make([]string, len(dateMatches))
	timeStrings := make([]string, len(dateMatches))

	for i, j := 0, 0; i < len(dateMatches); i++ {
		date := period[dateMatches[i][0]:dateMatches[i][1]]
		if j < len(timeMatches) && timeMatches[j][0] == dateMatches[i][1] {
			t := period[timeMatches[j][0]:timeMatches[j][1]]
			j++

			timeStrings[i] = fmt.Sprint(date, t)
			formats[i] = fmt.Sprint(SimpleDateFormatV2, "T", SimpleTimeFormat)
		} else {
			timeStrings[i] = date
			formats[i] = SimpleDateFormatV2
		}
	}

	times := [2]time.Time{}
	for i, timeString := range timeStrings {
		var err error
		times[i], err = time.Parse(formats[i], timeString)
		if err != nil {
			return TimePeriod{}, err
		}
	}

	timePeriod := TimePeriod{
		Start: times[0],
	}

	if !times[1].IsZero() {
		timePeriod.End = times[1]
	} else {
		timePeriod.End = time.Now()
	}

	return timePeriod, nil
}

// Parse takes a string in one of the following formats and returns a new TimePeriod, and optionally, an error
//
// Deprecated: Please use ParseTimePeriod instead
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

// In returns a copy of TimePeriod tp with both start and end times' location set to the specified location
func (tp TimePeriod) In(location *time.Location) TimePeriod {
	return TimePeriod{
		Start: tp.Start.In(location),
		End:   tp.End.In(location),
	}
}

// UTC returns a copy of TimePeriod tp with both start and end times' location set to UTC
func (tp TimePeriod) UTC() TimePeriod {
	return tp.In(time.UTC)
}

// Length returns the length of the period as a time.Duration value
func (tp TimePeriod) Length() time.Duration {
	return tp.End.Sub(tp.Start)
}

// Since returns the amount of time elapsed since the end of another TimePeriod as a time.Duration value
func (tp TimePeriod) Since(other TimePeriod) time.Duration {
	return tp.Start.Sub(other.End)
}

// Format returns the string representation of this timePeriod in the given format
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
	layout := fmt.Sprint(SimpleDateFormatV2, "T", SimpleTimeFormat)
	return tp.Format(layout)
}

// NewTimePeriod returns a TimePeriod starting at the given time and ending at the given time plus the given duration
func NewTimePeriod(start time.Time, period time.Duration) TimePeriod {
	return TimePeriod{
		Start: start,
		End:   start.Add(period),
	}
}
