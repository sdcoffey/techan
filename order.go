package techan

import (
	"time"

	"github.com/sdcoffey/big"
)

// OrderSide is a simple enumeration representing the side of an Order (buy or sell)
type OrderSide int

// BUY and SELL enumerations
const (
	BUY OrderSide = iota
	SELL
)

func (os OrderSide) String() string {
	switch os {
	case BUY:
		return "BUY"
	case SELL:
		return "SELL"
	default:
		return "UNKNOWN"
	}
}

// Order represents a trade execution (buy or sell) with associated metadata.
type Order struct {
	Side          OrderSide
	Security      string
	Price         big.Decimal
	Amount        big.Decimal
	ExecutionTime time.Time
}
