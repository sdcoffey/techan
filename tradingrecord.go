package talib4g

import (
	"time"
)

type TradingRecord struct {
	Trades       []*Position
	currentTrade *Position
}

func NewTradingRecord() (t *TradingRecord) {
	t = new(TradingRecord)
	t.Trades = make([]*Position, 0)
	t.currentTrade = newPosition()
	return t
}

func (this *TradingRecord) CurrentTrade() *Position {
	return this.currentTrade
}

func (this *TradingRecord) LastTrade() *Position {
	if len(this.Trades) == 0 {
		return nil
	}

	return this.Trades[len(this.Trades)-1]
}

func (this *TradingRecord) Enter(price, amount Money, time time.Time) {
	order := NewOrder(BUY)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time

	this.operate(order)
}

func (this *TradingRecord) Exit(price, amount Money, time time.Time) {
	order := NewOrder(SELL)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time

	this.operate(order)
}

func (this *TradingRecord) operate(order *order) {
	if this.currentTrade.IsOpen() {
		if order.ExecutionTime.Before(this.CurrentTrade().EntranceOrder().ExecutionTime) {
			return
		}

		this.currentTrade.Exit(order)
		this.Trades = append(this.Trades, this.currentTrade)
		this.currentTrade = newPosition()
	} else if this.currentTrade.IsNew() {
		if this.LastTrade() != nil && order.ExecutionTime.Before(this.LastTrade().ExitOrder().ExecutionTime) {
			return
		}

		this.currentTrade.Enter(order)
	}
}
