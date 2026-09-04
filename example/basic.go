package example

import (
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

// BasicEma is an example of how to create a basic Exponential moving average indicator
// based on the close prices of a timeseries from your exchange of choice.
func BasicEma() techan.Indicator {
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1609459200", "99", "100", "101", "98", "12"},
		{"1609545600", "100", "102", "103", "99", "15"},
		{"1609632000", "102", "101", "104", "100", "11"},
		{"1609718400", "101", "104", "105", "100", "18"},
	}

	for _, datum := range dataset {
		start, err := strconv.ParseInt(datum[0], 10, 64)
		if err != nil {
			panic(err)
		}
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[2])
		candle.MaxPrice = big.NewFromString(datum[3])
		candle.MinPrice = big.NewFromString(datum[4])
		candle.Volume = big.NewFromString(datum[5])

		if !series.AddCandle(candle) {
			panic("candles must be in chronological order")
		}
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, 3) // Three observations per window

	return movingAverage
}
