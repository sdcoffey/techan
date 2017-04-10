package talib4g

import (
	"time"

	"github.com/shopspring/decimal"
)

type TradingRecord struct {
	Trades       []*Trade
	currentTrade *Trade
}

func NewTradingRecord() (t *TradingRecord) {
	t = new(TradingRecord)
	t.currentTrade = newTrade()
	return t
}

func (this *TradingRecord) CurrentTrade() *Trade {
	return this.currentTrade
}

func (this *TradingRecord) Enter(price, amount decimal.Decimal, time time.Time) {
	order := NewOrder(BUY)
	order.Amount = amount
	order.Price = price

	this.operate(order)
}

func (this *TradingRecord) Exit(price, amount decimal.Decimal, time time.Time) {
	order := NewOrder(SELL)
	order.Amount = amount
	order.Price = price

	this.operate(order)
}

func (this *TradingRecord) operate(order *Order) {
	if this.currentTrade.IsOpen() {
		this.currentTrade.Exit(order)
		this.Trades = append(this.Trades, this.currentTrade)
		this.currentTrade = newTrade()
	} else if this.currentTrade.IsNew() {
		this.currentTrade.Enter(order)
	}
}
