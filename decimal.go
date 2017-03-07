package talib4g

import "github.com/shopspring/decimal"

func NewDecimal(value int) decimal.Decimal {
	return decimal.New(int64(value), 0)
}

var (
	ZERO  decimal.Decimal = NewDecimal(0)
	ONE   decimal.Decimal = NewDecimal(1)
	TWO   decimal.Decimal = NewDecimal(2)
	THREE decimal.Decimal = NewDecimal(3)
	FOUR  decimal.Decimal = NewDecimal(4)
	FIVE  decimal.Decimal = NewDecimal(5)
	SIX   decimal.Decimal = NewDecimal(6)
	SEVEN decimal.Decimal = NewDecimal(7)
	EIGHT decimal.Decimal = NewDecimal(8)
	NINE  decimal.Decimal = NewDecimal(9)
	TEN   decimal.Decimal = NewDecimal(10)
)
