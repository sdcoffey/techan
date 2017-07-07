package talib4g

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMoney(t *testing.T) {
	t.Run("Returns error if zero length string", func(t *testing.T) {
		_, err := ParseMoney("")
		assert.EqualError(t, err, "cannot parse string ")
	})

	t.Run("Returns error if symbol no symbol present", func(t *testing.T) {
		_, err := ParseMoney("10")
		assert.EqualError(t, err, "cannot parse symbol 1")
	})

	t.Run("Returns error if symbol not recognized", func(t *testing.T) {
		_, err := ParseMoney("™10")
		assert.EqualError(t, err, "cannot parse symbol ™")
	})

	t.Run("Returns an error if decimal is unparseable", func(t *testing.T) {
		_, err := ParseMoney("$dhjsk")
		assert.Error(t, err)
	})

	t.Run("Parses well-formed money successfully", func(t *testing.T) {
		m, err := ParseMoney("€4.21")
		assert.NoError(t, err)

		assert.EqualValues(t, 4.21, m.Float())
		assert.EqualValues(t, "EUR", m.Currency.String())
	})

	t.Run("Parses all zeroes - zero", func(t *testing.T) {
		m, err := ParseMoney("€0")
		assert.NoError(t, err)

		assert.EqualValues(t, 0, m.Float())
		assert.EqualValues(t, "EUR", m.Currency.String())
	})

	t.Run("Parses all zeroes - multiple", func(t *testing.T) {
		m, err := ParseMoney("€0000")
		assert.NoError(t, err)

		assert.EqualValues(t, 0, m.Float())
		assert.EqualValues(t, "EUR", m.Currency.String())
	})

	t.Run("Handles non-total amount of zeroes in the decimal place < ", func(t *testing.T) {
		m, err := ParseMoney("$4.1")
		assert.NoError(t, err)

		assert.EqualValues(t, 4.1, m.Float())
	})

	t.Run("Handles non-total amount of zeroes in the decimal place > ", func(t *testing.T) {
		m, err := ParseMoney("$4.10000000")
		assert.NoError(t, err)

		assert.EqualValues(t, 4.1, m.Float())
	})
}

func TestMoney_New(t *testing.T) {
	t.Run("New with float", func(t *testing.T) {
		money := NM(1.50, USD)
		assert.EqualValues(t, 150, money.Value())
	})

	t.Run("New with int", func(t *testing.T) {
		money := NMI(1, USD)
		assert.EqualValues(t, 100, money.Value())
	})

	t.Run("New with raw", func(t *testing.T) {
		money := NM(1, USD)
		assert.EqualValues(t, 100, money.Value())
	})
}

func TestMoney_Add(t *testing.T) {
	money := NM(1.50, USD)

	another := NM(1, USD)
	money = money.A(another)

	assert.EqualValues(t, 250, money.Value())
}

func TestMoney_Sub(t *testing.T) {
	money := NM(1.50, USD)

	another := NM(1, USD)
	money = money.S(another)

	assert.EqualValues(t, 50, money.Value())

	money = money.S(NM(2, USD))

	assert.EqualValues(t, -150, money.Value())
}

func TestMoney_Mul(t *testing.T) {
	money := NM(10, USD)
	money = money.M(NM(10, USD))

	assert.EqualValues(t, NM(100, USD).Value(), money.Value())
}

func TestMoney_Div(t *testing.T) {
	money := NM(10, USD)
	money = money.D(NM(12, USD))

	assert.EqualValues(t, NM(.83, USD).Value(), money.Value())
}

func TestMoney_Convert(t *testing.T) {
	btc := NM(10, BTC)
	usdWorth := btc.Convert(NM(100, USD))

	assert.EqualValues(t, NM(1000, USD).Value(), usdWorth.Value())

	usd := NM(1, USD)
	eurWorth := usd.Convert(NM(.5, EUR))

	assert.EqualValues(t, NM(0.5, EUR).Value(), eurWorth.Value())
}

func TestMoney_String(t *testing.T) {
	money := NM(10.38, USD)
	assert.Equal(t, "10.38", money.String())
}

func TestMoney_FormalString(t *testing.T) {
	money := NM(10.38, USD)
	assert.Equal(t, "$10.38", money.FormalString())
}

func TestMoney_Float(t *testing.T) {
	money := NM(10.38, BTC)
	assert.EqualValues(t, 10.38, money.Float())
}

func TestMoney_Abs(t *testing.T) {
	t.Run("Is negative", func(t *testing.T) {
		money := NM(-10, USD)
		assert.EqualValues(t, 10, money.Abs().Float())
	})

	t.Run("Is positive", func(t *testing.T) {
		money := NM(10, USD)
		assert.EqualValues(t, 10, money.Abs().Float())
	})

	t.Run("Is zero", func(t *testing.T) {
		money := NM(0, USD)
		assert.EqualValues(t, 0, money.Abs().Float())
	})
}

func TestMoney_Frac(t *testing.T) {
	t.Run("Less than 1", func(t *testing.T) {
		money := NM(10, USD)
		money = money.Frac(.5)

		assert.EqualValues(t, 5, money.Float())

		money = NM(.25, USD)
		money = money.Frac(.5)

		assert.EqualValues(t, .12, money.Float())
	})

	t.Run("Greater than 1", func(t *testing.T) {
		money := NM(10, USD)
		money = money.Frac(2)

		assert.EqualValues(t, 20, money.Float())
	})
}

func TestMoney_Neg(t *testing.T) {
	money := NM(-10, USD)
	assert.EqualValues(t, 10, money.Neg().Float())

	money = NM(10, USD)
	assert.EqualValues(t, -10, money.Neg().Float())
}

func TestMoney_MarshalJSON(t *testing.T) {
	c := NM(10, USD)
	jsonVal, err := json.Marshal(c)
	assert.NoError(t, err)

	assert.EqualValues(t, `{"Value":1000,"Currency":"USD"}`, string(jsonVal))
}

func TestMoney_UnmarshalJSON(t *testing.T) {
	value := `{"Value":1000,"Currency":"USD"}`
	var money Money

	err := json.Unmarshal([]byte(value), &money)
	assert.NoError(t, err)

	assert.EqualValues(t, 10, money.Float())
	assert.EqualValues(t, "USD", money.Currency.String())
}

func TestCurrency_MarshalJSON(t *testing.T) {
	c := USD
	jsonValue, err := json.Marshal(c)
	assert.NoError(t, err)
	assert.EqualValues(t, `{"Label":"USD"}`, jsonValue)
}

func TestCurrency_UnmarshalJSON(t *testing.T) {
	marshaled := `{"Label":"USD"}`
	var currency Currency

	err := json.Unmarshal([]byte(marshaled), &currency)
	assert.NoError(t, err)

	assert.EqualValues(t, "USD", currency.String())
}

func BenchmarkAdd(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money = money.A(NM(1, USD))
	}
}

func BenchmarkSub(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money = money.S(NM(1, USD))
	}
}

func BenchmarkMul(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money = money.M(NM(100, USD))
	}
}

func BenchmarkDiv(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money = money.D(NM(2, USD))
	}
}
