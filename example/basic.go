package example

import (
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/talib4g"
)

func BasicEma() talib4g.Indicator {
	series := talib4g.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := talib4g.NewTimePeriodD(time.Unix(start, 0), time.Hour*24)

		candle := talib4g.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[2])
		candle.MaxPrice = big.NewFromString(datum[3])
		candle.MinPrice = big.NewFromString(datum[4])

		series.AddCandle(candle)
	}

	closePrices := talib4g.NewClosePriceIndicator(series)
	movingAverage := talib4g.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

	return movingAverage
}
