package talib4g

import (
	"github.com/shopspring/decimal"
	"time"
)

type Tick struct {
	Period     time.Duration
	BeginTime  time.Time
	EndTime    time.Time
	OpenPrice  decimal.Decimal
	ClosePrice decimal.Decimal
	MaxPrice   decimal.Decimal
	MinPrice   decimal.Decimal
	Amount     decimal.Decimal
	Volume     decimal.Decimal
	TradeCount uint
}

func NewTick(period time.Duration, endTime time.Time) (t *Tick) {
	t = new(Tick)

	t.Period = period
	t.EndTime = endTime
	t.BeginTime = t.EndTime.Add(-t.Period)

	return t
}

func (this *Tick) AddTrade(tradeAmount, tradePrice decimal.Decimal) {
	if this.OpenPrice == decimal.Zero {
		this.OpenPrice = tradePrice
	}
	this.ClosePrice = tradePrice

	if this.MaxPrice == decimal.Zero {
		this.MaxPrice = tradePrice
	} else if tradePrice.Cmp(this.MaxPrice) > 0 {
		this.MaxPrice = tradePrice
	}

	if this.MinPrice == decimal.Zero {
		this.MinPrice = tradePrice
	} else if tradePrice.Cmp(this.MinPrice) < 0 {
		this.MinPrice = tradePrice
	}

	this.Amount = this.Amount.Add(tradeAmount)
	this.Volume = this.Volume.Add(tradeAmount.Mul(tradePrice))

	this.TradeCount++
}
