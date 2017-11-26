package talib4g

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var candleIndex int

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

	return mockTimeSeries(vals...)
}

func mockTimeSeries(values ...float64) *TimeSeries {
	ts := NewTimeSeries()
	for _, val := range values {
		candle := NewCandle(NewTimePeriodD(time.Unix(int64(candleIndex), 0), time.Second))
		candle.ClosePrice = big.NewDecimal(val)
		candle.Volume = big.NewDecimal(val)

		ts.AddCandle(candle)

		candleIndex++
	}

	return ts
}

func decimalEquals(t *testing.T, expected float64, actual big.Decimal) {
	assert.Equal(t, fmt.Sprintf("%.4f", expected), fmt.Sprintf("%.4f", actual.Float()))
}
