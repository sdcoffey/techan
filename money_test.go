package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		money := nmr(100, USD)
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
	usd := NM(1000, USD)
	btc := usd.Convert(NM(100, BTC))

	assert.EqualValues(t, NM(10, BTC).Value(), btc.Value())

	usd = NM(1, USD)
	eur := usd.Convert(NM(1.2, EUR))

	assert.EqualValues(t, NM(.83, EUR).Value(), eur.Value())
}

func TestMoney_String(t *testing.T) {
	money := NM(10.38, USD)
	assert.EqualValues(t, "10.38", money.String())
}

func TestMoney_Float(t *testing.T) {
	money := NM(10.38, BTC)
	assert.EqualValues(t, 10.38, money.Float())
}

func BenchmarkAdd(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money.A(NM(1, USD))
	}
}

func BenchmarkSub(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money.S(NM(1, USD))
	}
}

func BenchmarkMul(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money.M(NM(100, USD))
	}
}

func BenchmarkDiv(b *testing.B) {
	money := NM(100, USD)

	for i := 0; i < b.N; i++ {
		money.D(NM(2, USD))
	}
}
