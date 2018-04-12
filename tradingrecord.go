package techan

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

// Operate takes an order and adds it to the current TradingRecord. It will only add the order if:
// - The current position is open and the passed order was executed after the entrance order
// - The current position is new and the passed order was executed after the last exit order
func (tr *TradingRecord) Operate(order Order) {
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
