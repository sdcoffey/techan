package techan

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var (
	candleIndex      int
	mockedTimeSeries = mockTimeSeriesFl(
		64.75, 63.79, 63.73,
		63.73, 63.55, 63.19,
		63.91, 63.85, 62.95,
		63.37, 61.33, 61.51)
)

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

func mockTimeSeriesOCHL(values ...[]float64) *TimeSeries {
	ts := NewTimeSeries()
	for i, ochl := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.OpenPrice = big.NewDecimal(ochl[0])
		candle.ClosePrice = big.NewDecimal(ochl[1])
		candle.MaxPrice = big.NewDecimal(ochl[2])
		candle.MinPrice = big.NewDecimal(ochl[3])
		candle.Volume = big.NewDecimal(float64(i))

		ts.AddCandle(candle)
	}

	return ts
}

func mockTimeSeriesOCHLV(values ...[]float64) *TimeSeries {
	ts := NewTimeSeries()
	for i, ochlv := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.OpenPrice = big.NewDecimal(ochlv[0])
		candle.ClosePrice = big.NewDecimal(ochlv[1])
		candle.MaxPrice = big.NewDecimal(ochlv[2])
		candle.MinPrice = big.NewDecimal(ochlv[3])
		candle.Volume = big.NewDecimal(ochlv[4])

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
		candle.MaxPrice = big.NewFromString(val).Add(big.ONE)
		candle.MinPrice = big.NewFromString(val).Sub(big.ONE)
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

func dump(indicator Indicator) (values []float64) {
	precision := 4.0
	m := math.Pow(10, precision)

	defer func() {
		recover()
	}()

	var index int
	for {
		values = append(values, math.Round(indicator.Calculate(index).Float()*m)/m)
		index++
	}

	return
}

func indicatorEquals(t *testing.T, expected []float64, indicator Indicator) {
	actualValues := dump(indicator)
	assert.EqualValues(t, expected, actualValues)
}
