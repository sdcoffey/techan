package test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/sdcoffey/talib4g"
	"github.com/shopspring/decimal"
	. "github.com/stretchr/testify/assert"
)

var tickIndex int

func RandomTimeSeries(size int) *TimeSeries {
	vals := make([]float64, size)
	for i := 0; i < size; i++ {
		vals[i] = rand.Float64() * 100
	}

	return MockTimeSeries(vals...)
}

func MockTick(closePrice float64) *Tick {
	t := NewTick(time.Second, time.Unix(int64(tickIndex), 0))
	t.ClosePrice = closePrice

	return t
}

func MockTimeSeries(values ...float64) *TimeSeries {
	ts := NewTimeSeries()
	for _, val := range values {
		tick := NewTick(time.Second, time.Unix(int64(tickIndex), 0))
		tick.ClosePrice = val

		ts.AddTick(tick)

		tickIndex++
	}

	return ts
}

func decimalEquals(t *testing.T, expected float64, actual float64) {
	d := decimal.NewFromFloat(expected)
	e := decimal.NewFromFloat(actual)
	Equal(t, d.StringFixed(4), e.StringFixed(4))
}
