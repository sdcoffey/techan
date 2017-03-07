package test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
	. "github.com/stretchr/testify/assert"
)

func RandomTimeSeries(size int) *TimeSeries {
	vals := make([]float64, size)
	for i := 0; i < size; i++ {
		vals[i] = rand.Float64() * 100
	}

	return MockTimeSeries(vals...)
}

func MockTimeSeries(values ...float64) *TimeSeries {
	ts := NewTimeSeries()
	for i, val := range values {
		tick := NewTick(time.Second, time.Unix(int64(i), 0))
		tick.ClosePrice = decimal.NewFromFloat(val)

		ts.AddTick(tick)
	}

	return ts
}

func decimalEquals(t *testing.T, expected float64, actual decimal.Decimal) {
	d := decimal.NewFromFloat(expected)
	Equal(t, d.StringFixed(4), actual.StringFixed(4))
}
