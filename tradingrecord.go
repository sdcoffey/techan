package talib4g

import (
	"time"

	"github.com/sdcoffey/big"
)

type TradingRecord struct {
	Trades          []*Position
	currentPosition *Position
}

func NewTradingRecord() (t *TradingRecord) {
	t = new(TradingRecord)
	t.Trades = make([]*Position, 0)
	t.currentPosition = new(Position)
	return t
}

func (this *TradingRecord) CurrentPosition() *Position {
	return this.currentPosition
}

func (this *TradingRecord) LastTrade() *Position {
	if len(this.Trades) == 0 {
		return nil
	}

	return this.Trades[len(this.Trades)-1]
}

func (this *TradingRecord) Enter(price, amount, feePercentage big.Decimal, security string, time time.Time) {
	order := NewOrder(BUY)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time
	order.Security = security
	order.FeePercentage = feePercentage

	this.operate(order)
}

func (this *TradingRecord) Exit(price, amount, feePercentage big.Decimal, security string, time time.Time) {
	order := NewOrder(SELL)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time
	order.Security = security
	order.FeePercentage = feePercentage

	this.operate(order)
}

func (this *TradingRecord) operate(order *Order) {
	if this.currentPosition.IsOpen() {
		if order.ExecutionTime.Before(this.CurrentPosition().EntranceOrder().ExecutionTime) {
			return
		}

		this.currentPosition.Exit(order)
		this.Trades = append(this.Trades, this.currentPosition)
		this.currentPosition = new(Position)
	} else if this.currentPosition.IsNew() {
		if this.LastTrade() != nil && order.ExecutionTime.Before(this.LastTrade().ExitOrder().ExecutionTime) {
			return
		}

		this.currentPosition.Enter(order)
	}
}
