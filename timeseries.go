package talib4g

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type TimeSeries struct {
	Ticks []*Tick
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

func (this *TimeSeries) Run(strategy Strategy) *TradingRecord {
	record := NewTradingRecord()

	for i, tick := range this.Ticks {
		if strategy.ShouldEnter(i, record) {
			order := NewOrder(BUY)
			order.Price = tick.ClosePrice
			order.Amount = decimal.NewFromFloat(0.01)
			order.ExecutionTime = tick.EndTime

			record.Enter(order.Price, order.Amount, order.ExecutionTime)
		} else if strategy.ShouldExit(i, record) {
			order := NewOrder(SELL)
			order.Amount = decimal.NewFromFloat(0.01)
			order.Price = tick.ClosePrice
			order.ExecutionTime = tick.EndTime

			record.Exit(order.Price, order.Amount, order.ExecutionTime)
		}
		if i%100 == 0 {
			println(i)
		}
	}

	return record
}
