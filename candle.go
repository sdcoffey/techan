package talib4g

import (
	"fmt"
	"time"
)

type Candle struct {
	Period     time.Duration
	BeginTime  time.Time
	EndTime    time.Time
	OpenPrice  float64 `json:",string"`
	ClosePrice float64 `json:",string"`
	MaxPrice   float64 `json:",string"`
	MinPrice   float64 `json:",string"`
	Volume     float64 `json:",string"`
	TradeCount uint
}

func NewCandle(period time.Duration, endTime time.Time) (c *Candle) {
	c = new(Candle)

	c.Period = period
	c.EndTime = endTime
	c.BeginTime = c.EndTime.Add(-c.Period)

	return c
}

func (this *Candle) AddTrade(tradeAmount, tradePrice float64) {
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

	this.Volume += tradeAmount
	this.TradeCount++
}

func (this *Candle) String() string {
	return fmt.Sprintf("%s \nOpen: %.5f\nClose: %.5f\nVolume: %.5f",
		this.EndTime.Format(time.Stamp),
		this.MaxPrice,
		this.MinPrice,
		this.Volume,
	)
}
