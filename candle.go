package talib4g

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type Candle struct {
	Period     TimePeriod
	OpenPrice  big.Decimal `json:",string"`
	ClosePrice big.Decimal `json:",string"`
	MaxPrice   big.Decimal `json:",string"`
	MinPrice   big.Decimal `json:",string"`
	Volume     big.Decimal `json:",string"`
	TradeCount uint
}

func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  big.ZERO,
		ClosePrice: big.ZERO,
		MaxPrice:   big.ZERO,
		MinPrice:   big.ZERO,
		Volume:     big.ZERO,
	}
}

func (c *Candle) AddTrade(tradeAmount, tradePrice big.Decimal) {
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
