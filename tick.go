package talib4g

import (
	"fmt"
	"time"
)

type Tick struct {
	Period     time.Duration
	BeginTime  time.Time
	EndTime    time.Time
	OpenPrice  float64 `json:",string"`
	ClosePrice float64 `json:",string"`
	MaxPrice   float64 `json:",string"`
	MinPrice   float64 `json:",string"`
	Amount     float64 `json:",string"`
	Volume     float64 `json:",string"`
	TradeCount uint
}

func NewTick(period time.Duration, endTime time.Time) (t *Tick) {
	t = new(Tick)

	t.Period = period
	t.EndTime = endTime
	t.BeginTime = t.EndTime.Add(-t.Period)

	return t
}

func (this *Tick) AddTrade(tradeAmount, tradePrice float64) {
	if this.OpenPrice == 0 {
		this.OpenPrice = tradePrice
	}
	this.ClosePrice = tradePrice

	if this.MaxPrice == 0 {
		this.MaxPrice = tradePrice
	} else if tradePrice > this.MaxPrice {
		this.MaxPrice = tradePrice
	}

	if this.MinPrice == 0 {
		this.MinPrice = tradePrice
	} else if tradePrice < this.MinPrice {
		this.MinPrice = tradePrice
	}

	this.Amount += tradeAmount
	this.Volume += (tradeAmount * tradePrice)

	this.TradeCount++
}

func (this *Tick) String() string {
	return fmt.Sprintf("Tick ending at %s - HL: %.5f/%.5f V: %.5f",
		this.EndTime.Format(time.Stamp),
		this.MaxPrice,
		this.MinPrice,
		this.Volume,
	)
}
