package techan

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"strconv"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var candleIndex int

func randomTimeSeries(size int) *TimeSeries {
	vals := make([]string, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		val := rand.Float64() * 100
		if i == 0 {
			vals[i] = fmt.Sprint(val)
		} else {
			last, _ := strconv.ParseFloat(vals[i-1], 64)
			if i%2 == 0 {
				vals[i] = fmt.Sprint(last + (val / 10))
			} else {
				vals[i] = fmt.Sprint(last - (val / 10))
			}
		}
	}

	return mockTimeSeries(vals...)
}

func mockTimeSeriesOCHL(values ...[]string) *TimeSeries {
	ts := NewTimeSeries()
	for i, ochl := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.OpenPrice = big.NewFromString(ochl[0])
		candle.ClosePrice = big.NewFromString(ochl[1])
		candle.MaxPrice = big.NewFromString(ochl[2])
		candle.MinPrice = big.NewFromString(ochl[3])
		candle.Volume = big.NewDecimal(float64(i))

		ts.AddCandle(candle)
	}

	return ts
}

func mockTimeSeries(values ...string) *TimeSeries {
	ts := NewTimeSeries()
	for _, val := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))
		candle.OpenPrice = big.NewFromString(val)
		candle.ClosePrice = big.NewFromString(val)
		candle.MaxPrice = big.NewFromString(val)
		candle.MinPrice = big.NewFromString(val)
		candle.Volume = big.NewFromString(val)

		ts.AddCandle(candle)

		candleIndex++
	}

	return ts
}

func mockTimeSeriesFl(values ...float64) *TimeSeries {
	strVals := make([]string, len(values))

	for i, val := range values {
		strVals[i] = fmt.Sprint(val)
	}

	return mockTimeSeries(strVals...)
}

func decimalEquals(t *testing.T, expected float64, actual big.Decimal) {
	assert.Equal(t, fmt.Sprintf("%.4f", expected), fmt.Sprintf("%.4f", actual.Float()))
}
