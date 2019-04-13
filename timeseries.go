package techan

import (
	"fmt"
)

// TimeSeries represents an array of candles
type TimeSeries struct {
	Candles []*Candle
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Candles = make([]*Candle, 0)

	return t
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *TimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		panic(fmt.Errorf("error adding Candle: candle cannot be nil"))
	}

	if ts.LastCandle() == nil || candle.Period.Since(ts.LastCandle().Period) >= 0 {
		ts.Candles = append(ts.Candles, candle)
		return true
	}

	return false
}

// LastCandle will return the lastCandle in this series, or nil if this series is empty
func (ts *TimeSeries) LastCandle() *Candle {
	if len(ts.Candles) > 0 {
		return ts.Candles[len(ts.Candles)-1]
	}

	return nil
}

// LastIndex will return the index of the last candle in this series
func (ts *TimeSeries) LastIndex() int {
	return len(ts.Candles) - 1
}
