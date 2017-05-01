package talib4g

import (
	"time"
)

type TradingRecord struct {
	Trades       []*Trade
	currentTrade *Trade
}

func NewTradingRecord() (t *TradingRecord) {
	t = new(TradingRecord)
	t.Trades = make([]*Trade, 0)
	t.currentTrade = newTrade()
	return t
}

func (this *TradingRecord) CurrentTrade() *Trade {
	return this.currentTrade
}

func (this *TradingRecord) Enter(price, amount float64, time time.Time) {
	order := NewOrder(BUY)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time

	this.operate(order)
}

func (this *TradingRecord) Exit(price, amount float64, time time.Time) {
	order := NewOrder(SELL)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time

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
