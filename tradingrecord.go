package talib4g

import (
	"time"

	"github.com/sdcoffey/big"
)

// TradingRecord is an object describing a series of trades made and a current position
type TradingRecord struct {
	Trades          []*Position
	currentPosition *Position
}

// NewTradingRecord returns a new TradingRecord
func NewTradingRecord() (t *TradingRecord) {
	t = new(TradingRecord)
	t.Trades = make([]*Position, 0)
	t.currentPosition = new(Position)
	return t
}

// CurrentPosition returns the current position in this record
func (tr *TradingRecord) CurrentPosition() *Position {
	return tr.currentPosition
}

// LastTrade returns the last trade executed in this record
func (tr *TradingRecord) LastTrade() *Position {
	if len(tr.Trades) == 0 {
		return nil
	}

	return tr.Trades[len(tr.Trades)-1]
}

// Enter records an entrance order in the current position of this trading record
func (tr *TradingRecord) Enter(price, amount, feePercentage big.Decimal, security string, time time.Time) {
	order := NewOrder(BUY)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time
	order.Security = security
	order.FeePercentage = feePercentage

	tr.operate(order)
}

// Exit records an exit order in the current position of this trading record
func (tr *TradingRecord) Exit(price, amount, feePercentage big.Decimal, security string, time time.Time) {
	order := NewOrder(SELL)
	order.Amount = amount
	order.Price = price
	order.ExecutionTime = time
	order.Security = security
	order.FeePercentage = feePercentage

	tr.operate(order)
}

func (tr *TradingRecord) operate(order *Order) {
	if tr.currentPosition.IsOpen() {
		if order.ExecutionTime.Before(tr.CurrentPosition().EntranceOrder().ExecutionTime) {
			return
		}

		tr.currentPosition.Exit(order)
		tr.Trades = append(tr.Trades, tr.currentPosition)
		tr.currentPosition = new(Position)
	} else if tr.currentPosition.IsNew() {
		if tr.LastTrade() != nil && order.ExecutionTime.Before(tr.LastTrade().ExitOrder().ExecutionTime) {
			return
		}

		tr.currentPosition.Enter(order)
	}
}
