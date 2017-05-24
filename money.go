package talib4g

import (
	"math"
	"strconv"
)

type Calculatable interface {
	Value() int
}

type Money struct {
	*Currency
	raw int
}

func NM(rawVal float64, currency *Currency) Money {
	return Money{currency, int(rawVal * float64(currency.multiplier))}
}

func NMI(rawVal int, currency *Currency) Money {
	return Money{currency, rawVal * currency.multiplier}
}

func nmr(rawVal int, currency *Currency) Money {
	return Money{currency, rawVal}
}

func (m Money) A(c Calculatable) Money {
	return nmr(m.raw+c.Value(), m.Currency)
}

func (m Money) S(other Calculatable) Money {
	return nmr(m.raw-other.Value(), m.Currency)
}

func (m Money) M(other Calculatable) Money {
	return NMI((m.raw/m.multiplier)*(other.Value()/m.multiplier), m.Currency)

}

func (m Money) D(other Calculatable) Money {
	return NMI((m.raw/m.multiplier)/(other.Value()/m.multiplier), m.Currency)
}

// Returns a money in currency other, at the given exchange rate
func (m Money) Convert(exchangeRate Money) Money {
	v := (m.raw / m.multiplier) * (exchangeRate.raw / exchangeRate.multiplier)
	return NMI(v, exchangeRate.Currency)
}

func (m Money) Value() int {
	return m.raw
}

func (m Money) String() string {
	return strconv.FormatFloat(float64(m.raw)/float64(m.multiplier), 'f', m.decimal, 64)
}

type Currency struct {
	label      string
	multiplier int
	decimal    int
}

func newCurrency(label string, decimalPlace int) *Currency {
	return &Currency{
		label:      label,
		decimal:    decimalPlace,
		multiplier: int(math.Pow(10, float64(decimalPlace))),
	}
}

var (
	USD *Currency = newCurrency("USD", 2)
	BTC *Currency = newCurrency("BTC", 8)
)
