package talib4g

import (
	"fmt"
)

type TimeSeries struct {
	Candles []*Candle
}

func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Candles = make([]*Candle, 0)

	return t
}

func (ts *TimeSeries) AddCandle(candle *Candle) {
	if candle == nil {
		panic(fmt.Errorf("Error adding Candle: candle cannot be nil"))
	}

	if ts.LastCandle() == nil || candle.EndTime.After(ts.LastCandle().EndTime) {
		ts.Candles = append(ts.Candles, candle)
	}
}

func (ts *TimeSeries) LastCandle() *Candle {
	if len(ts.Candles) > 0 {
		return ts.Candles[len(ts.Candles)-1]
	}

	return nil
}

func (ts *TimeSeries) LastIndex() int {
	return len(ts.Candles) - 1
}
