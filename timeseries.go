package talib4g

import (
	"fmt"
)

type TimeSeries struct {
	BeginIndex int
	EndIndex   int
	Ticks      []*Tick
}

func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Ticks = make([]*Tick, 0)

	return t
}

func (this *TimeSeries) AddTick(tick *Tick) {
	if tick == nil {
		panic(fmt.Errorf("Error adding Tick: Tick cannot be nil"))
	}

	if len(this.Ticks) > 0 && !tick.EndTime.After(this.LastTick().EndTime) {
		panic(fmt.Errorf("Error adding Tick: cannot have end Time < series end time"))
	}

	this.Ticks = append(this.Ticks, tick)
}

func (this *TimeSeries) LastTick() *Tick {
	if len(this.Ticks) > 0 {
		return this.Ticks[len(this.Ticks)-1]
	}

	return nil
}
