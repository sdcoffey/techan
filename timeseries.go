package talib4g

import (
	"fmt"
	"time"
)

type TimeSeries struct {
	Ticks []*Tick
}

func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Ticks = make([]*Tick, 0)

	return t
}

func (this TimeSeries) TimeXValues() []time.Time {
	times := make([]time.Time, len(this.Ticks))
	for i, tick := range this.Ticks {
		times[i] = tick.EndTime
	}

	return times
}

func (this *TimeSeries) AddTick(tick *Tick) {
	if tick == nil {
		panic(fmt.Errorf("Error adding Tick: Tick cannot be nil"))
	}

	if this.LastTick() == nil || tick.EndTime.After(this.LastTick().EndTime) {
		this.Ticks = append(this.Ticks, tick)
	}
}

func (this *TimeSeries) LastTick() *Tick {
	if len(this.Ticks) > 0 {
		return this.Ticks[len(this.Ticks)-1]
	}

	return nil
}

func (this *TimeSeries) Run(strategy Strategy, startingAmount float64) *TradingRecord {
	record := NewTradingRecord()

	var openPositionAmount float64
	for i, tick := range this.Ticks {
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
