package talib4g

import (
	"math/big"
)

var (
	flZero = *big.NewFloat(0)

	ZERO = NewDecimal(0)
	ONE  = NewDecimal(1)
	TEN  = NewDecimal(10)
)

type Decimal struct {
	fl *big.Float
}

func NewDecimal(fl float64) Decimal {
	return Decimal{big.NewFloat(fl)}
}

func NewDecimalFromString(fl string) Decimal {
	bfl := new(big.Float)
	bfl.UnmarshalText([]byte(fl))
	return Decimal{bfl}
}

func (d Decimal) Add(other Decimal) Decimal {
	return Decimal{d.cpy().Add(d.fl, other.fl)}
}

func (d Decimal) Sub(other Decimal) Decimal {
	return Decimal{d.cpy().Sub(d.fl, other.fl)}
}

func (d Decimal) Mul(other Decimal) Decimal {
	return Decimal{d.cpy().Mul(d.fl, other.fl)}
}

func (d Decimal) Div(other Decimal) Decimal {
	return Decimal{d.cpy().Quo(d.fl, other.fl)}
}

func (d Decimal) Frac(f float64) Decimal {
	return d.Mul(NewDecimal(f))
}

func (d Decimal) Neg() Decimal {
	return d.Mul(NewDecimal(-1))
}

func (d Decimal) Abs() Decimal {
	if d.LT(ZERO) {
		return d.Mul(ONE.Neg())
	}

	return d
}

func (d Decimal) EQ(other Decimal) bool {
	return d.Cmp(other) == 0
}

func (d Decimal) LT(other Decimal) bool {
	return d.Cmp(other) < 0
}

func (d Decimal) LTE(other Decimal) bool {
	return d.Cmp(other) <= 0
}

func (d Decimal) GT(other Decimal) bool {
	return d.Cmp(other) > 0
}

func (d Decimal) GTE(other Decimal) bool {
	return d.Cmp(other) >= 0
}

func (d Decimal) Cmp(other Decimal) int {
	return d.fl.Cmp(other.fl)
}

func (d Decimal) Float() float64 {
	f, _ := d.fl.Float64()
	return f
}

func (d Decimal) Zero() bool {
	return d.fl.Cmp(&flZero) == 0
}

func (d Decimal) String() string {
	return d.fl.String()
}

func (d Decimal) cpy() *big.Float {
	cpy := new(big.Float)
	return cpy.Copy(d.fl)
}
