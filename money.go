package talib4g

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Calculatable interface {
	Value() int64
}

type Money struct {
	*Currency
	raw int64
}

func NS(value int) Money {
	return Money{security, int64(value)}
}

func NM(rawVal float64, currency *Currency) Money {
	return Money{currency, int64(rawVal * float64(currency.multiplier))}
}

func NMI(rawVal int64, currency *Currency) Money {
	return Money{currency, int64(rawVal) * currency.multiplier}
}

func nmr(rawVal int64, currency *Currency) Money {
	return Money{currency, rawVal}
}

func (m Money) A(c Calculatable) Money {
	return nmr(m.raw+c.Value(), m.Currency)
}

func (m Money) S(other Calculatable) Money {
	return nmr(m.raw-other.Value(), m.Currency)
}

func (m Money) M(other Calculatable) Money {
	lhs := m.raw / m.multiplier
	val := other.Value()
	rhs := val / m.multiplier
	return NMI(lhs*rhs, m.Currency)
}

func (m Money) D(other Calculatable) Money {
	return NM((float64(m.raw)/float64(m.multiplier))/(float64(other.Value())/float64(m.multiplier)), m.Currency)
}

func (m Money) GT(other Money) bool {
	return m.cmp(other) > 0
}

func (m Money) LT(other Money) bool {
	return m.cmp(other) < 0
}

func (m Money) EQ(other Money) bool {
	return m.cmp(other) == 0
}

func (m Money) Zero() bool {
	return m.raw == 0
}

func (m Money) Abs() Money {
	if m.raw < 0 {
		return Money{m.Currency, -m.raw}
	}

	return m
}

func (m Money) Neg() Money {
	return Money{m.Currency, -m.raw}
}

func (m Money) Frac(fraction float64) Money {
	return Money{m.Currency, int64(float64(m.raw) * fraction)}
}

// Returns a money in currency other, at the given exchange rate
func (m Money) Convert(exchangeRate Money) Money {
	return NM(m.Float()/(m.Float()/exchangeRate.Float())*m.Float(), exchangeRate.Currency)
}

func (m Money) Value() int64 {
	return m.raw
}

func (m Money) String() string {
	if m.Currency == nil {
		return "0"
	} else {
		return fmt.Sprintf("%s %s", m.Currency.label, strconv.FormatFloat(m.Float(), 'f', int(math.Log10(float64(m.multiplier))), 64))
	}
}

func (m Money) Float() float64 {
	return float64(m.raw) / float64(m.multiplier)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"Value":%d,"Currency":"%s"}`, m.raw, string(m.Currency.label))), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	split := bytes.Split(b, []byte(","))
	rawStr := split[0][9:]
	currencyStr := split[1][12:15]

	var err error
	m.raw, err = strconv.ParseInt(string(rawStr), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Value -> %s", err)
	}

	m.Currency = CurrencyForName(string(currencyStr))

	return nil
}

func (m Money) cmp(other Money) int {
	if m.Currency != other.Currency {
		panic(fmt.Errorf("Cannot compare two moneys of different currency"))
	}

	if m.raw == other.raw {
		return 0
	} else if m.raw < other.raw {
		return -1
	}
	return 1
}

type Currency struct {
	label      string
	multiplier int64
}

func (c *Currency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"Label":"%s"}`, c.label)), nil
}

func (c *Currency) UnmarshalJSON(b []byte) error {
	curr := CurrencyForName(string(b[10:13]))
	if curr == nil {
		return fmt.Errorf("No such currency: %s", string(b))
	}

	c.label = curr.label
	c.multiplier = curr.multiplier

	return nil
}

func (c *Currency) String() string {
	return c.label
}

func newCurrency(label string, decimalPlace int) *Currency {
	return &Currency{
		label:      label,
		multiplier: int64(math.Pow(10, float64(decimalPlace))),
	}
}

func CurrencyForName(name string) *Currency {
	switch strings.ToUpper(name) {
	case USD.label:
		return USD
	case EUR.label:
		return EUR
	case BTC.label:
		return BTC
	case ETH.label:
		return ETH
	}

	return nil
}

var (
	security *Currency = newCurrency("SEC", 0)
	USD      *Currency = newCurrency("USD", 2)
	EUR      *Currency = newCurrency("EUR", 2)
	GBP      *Currency = newCurrency("GBP", 2)
	BTC      *Currency = newCurrency("BTC", 8)
	ETH      *Currency = newCurrency("ETH", 18)
)
