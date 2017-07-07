package talib4g

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type Calculatable interface {
	Value() int64
}

type Money struct {
	*Currency
	raw big.Int
}

func ParseMoneyAs(moneyString string, currency *Currency) (Money, error) {
	return ParseMoney(currency.Symbol() + moneyString)
}

var zerosReg *regexp.Regexp = regexp.MustCompile("^0+$")

func ParseMoney(moneyString string) (Money, error) {
	if len(moneyString) == 0 {
		return Money{}, fmt.Errorf("cannot parse string %s", moneyString)
	}

	symbol := []rune(moneyString)[0]
	currency := CurrencyForSymbol(symbol)
	if currency == nil {
		return Money{}, fmt.Errorf("cannot parse symbol %s", string(symbol))
	}

	moneyString = strings.Split(moneyString, string(symbol))[1]

	if zerosReg.MatchString(moneyString) {
		return NM(0, currency), nil
	}

	dotIndex := strings.Index(moneyString, ".")
	if dotIndex < 0 {
		return Money{}, fmt.Errorf("Money string not well formed: %s", moneyString)
	}

	beforeDot, bdErr := strconv.ParseInt(string(moneyString[:dotIndex]), 10, 64)

	afterDotStr := string(moneyString[dotIndex+1:])
	if len(afterDotStr) > int(math.Log10(float64(currency.multiplier))) {
		afterDotStr = string(afterDotStr[:int(math.Log10(float64(currency.multiplier)))])
	} else {
		for len(afterDotStr) < int(math.Log10(float64(currency.multiplier))) {
			afterDotStr += "0"
		}
	}
	afterDot, adErr := strconv.ParseInt(afterDotStr, 10, 64)

	if bdErr != nil || adErr != nil {
		return Money{}, fmt.Errorf("Could not parse decimal %s -> %s", moneyString[1:], bdErr.Error()+adErr.Error())
	}

	rawVal := big.NewInt(beforeDot*currency.multiplier + afterDot)

	return nmr(rawVal, currency), nil
}

func NS(value int) Money {
	return Money{security, *big.NewInt(int64(value))}
}

func NM(rawVal float64, currency *Currency) Money {
	return Money{currency, *big.NewInt(int64(rawVal * float64(currency.multiplier)))}
}

func NMI(rawVal int64, currency *Currency) Money {
	return Money{currency, *big.NewInt(int64(rawVal) * currency.multiplier)}
}

func nmr(rawVal *big.Int, currency *Currency) Money {
	return Money{currency, *rawVal}
}

func (m Money) A(c Calculatable) Money {
	return nmr(m.raw.Add(&m.raw, big.NewInt(c.Value())), m.Currency)
}

func (m Money) S(c Calculatable) Money {
	return nmr(m.raw.Sub(&m.raw, big.NewInt(c.Value())), m.Currency)
}

func (m Money) M(other Calculatable) Money {
	lhs := m.raw.Div(&m.raw, big.NewInt(m.multiplier))
	rhs := big.NewInt(other.Value())
	rhs = rhs.Div(rhs, big.NewInt(m.multiplier))
	return NMI(rhs.Mul(rhs, lhs).Int64(), m.Currency)
}

func (m Money) D(other Calculatable) Money {
	return NM((float64(m.raw.Int64())/float64(m.multiplier))/(float64(other.Value())/float64(m.multiplier)), m.Currency)
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
	return m.raw.Int64() == 0
}

func (m Money) Abs() Money {
	if m.raw.Int64() < 0 {
		return Money{m.Currency, *m.raw.Neg(&m.raw)}
	}

	return m
}

func (m Money) Neg() Money {
	return Money{m.Currency, *m.raw.Neg(&m.raw)}
}

func (m Money) Frac(fraction float64) Money {
	return Money{m.Currency, *big.NewInt(int64(float64(m.raw.Int64()) * fraction))}
}

// Returns a money in currency other, at the given exchange rate
func (m Money) Convert(exchangeRate Money) Money {
	return NM(m.Float()/(m.Float()/exchangeRate.Float())*m.Float(), exchangeRate.Currency)
}

func (m Money) Value() int64 {
	return m.raw.Int64()
}

func (m Money) String() string {
	if m.Currency == nil {
		return "0"
	} else {
		return strconv.FormatFloat(m.Float(), 'f', int(math.Log10(float64(m.multiplier))), 64)
	}
}

func (m Money) FormalString() string {
	if m.Currency == nil {
		return "0"
	} else {
		return fmt.Sprintf("%s%s", string(m.Currency.symbol), m.String())
	}
}

func (m Money) Float() float64 {
	return float64(m.raw.Int64()) / float64(m.multiplier)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"Value":%d,"Currency":"%s"}`, m.raw.Int64(), string(m.Currency.label))), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	split := bytes.Split(b, []byte(","))
	rawStr := split[0][9:]
	currencyStr := split[1][12:15]

	r, err := strconv.ParseInt(string(rawStr), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Value -> %s", err)
	} else {
		m.raw = *big.NewInt(r)
	}

	m.Currency = CurrencyForName(string(currencyStr))

	return nil
}

func (m Money) cmp(other Money) int {
	if m.Currency != other.Currency {
		panic(fmt.Errorf("Cannot compare two moneys of different currency"))
	}

	if m.raw.Cmp(&other.raw) == 0 {
		return 0
	} else if m.raw.Cmp(&other.raw) < 0 {
		return -1
	}
	return 1
}

type Currency struct {
	label      string
	symbol     rune
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

func (c *Currency) Symbol() string {
	return string(c.symbol)
}

func newCurrency(label string, symbol rune, decimalPlace int) *Currency {
	return &Currency{
		label:      label,
		symbol:     symbol,
		multiplier: int64(math.Pow(10, float64(decimalPlace))),
	}
}

func CurrencyForName(name string) *Currency {
	switch strings.ToUpper(name) {
	case USD.label:
		return USD
	case EUR.label:
		return EUR
	case GBP.label:
		return GBP
	case CAD.label:
		return CAD
	case BTC.label:
		return BTC
	case ETH.label:
		return ETH
	case LTC.label:
		return LTC
	}

	return nil
}

func CurrencyForSymbol(symbol rune) *Currency {
	switch symbol {
	case USD.symbol:
		return USD
	case EUR.symbol:
		return EUR
	case GBP.symbol:
		return GBP
	case CAD.symbol:
		return CAD
	case BTC.symbol:
		return BTC
	case ETH.symbol:
		return ETH
	case LTC.symbol:
		return LTC
	}

	return nil
}

var (
	security *Currency = newCurrency("SEC", 'S', 0)
	USD      *Currency = newCurrency("USD", '$', 2)
	EUR      *Currency = newCurrency("EUR", '€', 2)
	GBP      *Currency = newCurrency("GBP", '£', 2)
	CAD      *Currency = newCurrency("CAD", 'C', 2)
	BTC      *Currency = newCurrency("BTC", 'B', 8)
	ETH      *Currency = newCurrency("ETH", 'E', 8)
	LTC      *Currency = newCurrency("LTC", 'L', 8)
)
