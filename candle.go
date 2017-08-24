package talib4g

import (
	"fmt"
)

type Candle struct {
	Period     TimePeriod
	OpenPrice  Decimal `json:",string"`
	ClosePrice Decimal `json:",string"`
	MaxPrice   Decimal `json:",string"`
	MinPrice   Decimal `json:",string"`
	Volume     Decimal `json:",string"`
	TradeCount uint
}

func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  ZERO,
		ClosePrice: ZERO,
		MaxPrice:   ZERO,
		MinPrice:   ZERO,
		Volume:     ZERO,
	}
}

func (c *Candle) AddTrade(tradeAmount, tradePrice Decimal) {
	if c.OpenPrice.Zero() {
		c.OpenPrice = tradePrice
	}
	c.ClosePrice = tradePrice

	if c.MaxPrice.Zero() {
		c.MaxPrice = tradePrice
	} else if tradePrice.GT(c.MaxPrice) {
		c.MaxPrice = tradePrice
	}

	if c.MinPrice.Zero() {
		c.MinPrice = tradePrice
	} else if tradePrice.LT(c.MinPrice) {
		c.MinPrice = tradePrice
	}

	if c.Volume.Zero() {
		c.Volume = tradeAmount
	} else {
		c.Volume = c.Volume.Add(tradeAmount)
	}

	c.TradeCount++
}

func (c *Candle) String() string {
	return fmt.Sprintf(
		`
	Time:	%s
	Open: 	%s
	Close: 	%s
	High: 	%s
	Low: 	%s
	Volume: %s
	`,
		c.Period,
		c.OpenPrice,
		c.ClosePrice,
		c.MaxPrice,
		c.MinPrice,
		c.Volume,
	)
}
