package talib4g

import "github.com/sdcoffey/big"

type Indicator interface {
	Calculate(int) big.Decimal
}
