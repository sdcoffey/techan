package talib4g

import (
	"fmt"
	"time"
)

type TimeSeries struct {
	Candles []*Candle
}

func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Candles = make([]*Candle, 0)

	return t
}

func (this TimeSeries) TimeXValues() []time.Time {
	times := make([]time.Time, len(this.Candles))
	for i, tick := range this.Candles {
		times[i] = tick.EndTime
	}

	return times
}

func (this *TimeSeries) AddCandle(tick *Candle) {
	if tick == nil {
		panic(fmt.Errorf("Error adding Tick: Tick cannot be nil"))
	}

	if this.LastCandle() == nil || tick.EndTime.After(this.LastCandle().EndTime) {
		this.Candles = append(this.Candles, tick)
	}
}

func (this *TimeSeries) LastCandle() *Candle {
	if len(this.Candles) > 0 {
		return this.Candles[len(this.Candles)-1]
	}

	return nil
}

func (this *TimeSeries) Run(strategy Strategy, startingAmount float64) *TradingRecord {
	record := NewTradingRecord()

	var openPositionAmount float64
	for i, tick := range this.Candles {
		if strategy.ShouldEnter(i, record) {
			openPositionAmount = startingAmount / tick.ClosePrice
			record.Enter(tick.ClosePrice, openPositionAmount, tick.EndTime)
		} else if strategy.ShouldExit(i, record) {
			record.Exit(tick.ClosePrice, openPositionAmount, tick.EndTime)
			openPositionAmount = 0
		}
	}

	return record
}
