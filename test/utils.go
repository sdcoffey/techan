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
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		val := rand.Float64() * 100
		if i == 0 {
			vals[i] = val
		} else {
			if i%2 == 0 {
				vals[i] = vals[i-1] + (val / 10)
			} else {
				vals[i] = vals[i-1] - (val / 10)
			}
		}
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
