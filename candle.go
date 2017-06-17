package talib4g

import (
	"fmt"
)

type Candle struct {
	Period     TimePeriod
	OpenPrice  Money `json:",string"`
	ClosePrice Money `json:",string"`
	MaxPrice   Money `json:",string"`
	MinPrice   Money `json:",string"`
	Volume     Money `json:",string"`
	TradeCount uint
}

func NewCandle(period TimePeriod) (c *Candle) {
	c = new(Candle)
	c.Period = period

	return c
}

func (this *Candle) AddTrade(tradeAmount, tradePrice Money) {
	if this.OpenPrice.Zero() {
		this.OpenPrice = tradePrice
	}
	this.ClosePrice = tradePrice

	if this.MaxPrice.Zero() {
		this.MaxPrice = tradePrice
	} else if tradePrice.GT(this.MaxPrice) {
		this.MaxPrice = tradePrice
	}

	if this.MinPrice.Zero() {
		this.MinPrice = tradePrice
	} else if tradePrice.LT(this.MinPrice) {
		this.MinPrice = tradePrice
	}

	if this.Volume.Zero() {
		this.Volume = tradeAmount
	} else {
		this.Volume = this.Volume.A(tradeAmount)
	}

	this.TradeCount++
}

func (this *Candle) String() string {
	return fmt.Sprintf(
		`
	Time:	%s
	Open: 	%s
	Close: 	%s
	High: 	%s
	Low: 	%s
	Volume: %s
	`,
		this.Period,
		this.OpenPrice,
		this.ClosePrice,
		this.MaxPrice,
		this.MinPrice,
		this.Volume,
	)
}
