package techan

import (
	"fmt"
)

// TimeSeries represents an array of candles
type TimeSeries interface {
	LastIndex() int
	FirstCandle() *Candle
	LastCandle() *Candle
	GetCandle(int) *Candle
	GetCandleData() []*Candle
	AddCandle(*Candle) bool
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries() TimeSeries {
	return new(BaseTimeSeries)
}

// BaseTimeSeries implements TimeSeries using in-memory slice
type BaseTimeSeries struct {
	Candles []*Candle
}

func (ts *BaseTimeSeries) GetCandle(idx int) *Candle {
	return ts.Candles[idx]
}

func (ts *BaseTimeSeries) GetCandleData() []*Candle {
	return ts.Candles
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *BaseTimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		panic(fmt.Errorf("error adding Candle: candle cannot be nil"))
	}

	if ts.LastCandle() == nil || candle.Period.Since(ts.LastCandle().Period) >= 0 {
		ts.Candles = append(ts.Candles, candle)
		return true
	}

	return false
}

// FirstCandle will return the firstCandle in this series, or nil if this series is empty
func (ts *BaseTimeSeries) FirstCandle() *Candle {
	if len(ts.Candles) > 0 {
		return ts.Candles[0]
	}

	return nil
}

// LastCandle will return the lastCandle in this series, or nil if this series is empty
func (ts *BaseTimeSeries) LastCandle() *Candle {
	if len(ts.Candles) > 0 {
		return ts.Candles[len(ts.Candles)-1]
	}

	return nil
}

// LastIndex will return the index of the last candle in this series
func (ts *BaseTimeSeries) LastIndex() int {
	return len(ts.Candles) - 1
}
