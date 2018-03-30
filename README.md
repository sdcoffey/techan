## Talib4g ![](https://travis-ci.org/sdcoffey/talib4g.svg?branch=master)

Talib4g is a library for technical analysis for Go! It provides a suite of tools and frameworks to analyze financial data and make trading decisions. 

## Features 
* Basic and advanced technical analysis indicators
* Profit and trade analysis
* Strategy building

### Installation
```sh
$ go get github.com/sdcoffey/talib4g
```

### Quickstart
```go
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

fmt.Println(movingAverage.Calculate(0).FormattedString(2))
```

### Creating trading strategies
```go
indicator := talib4g.NewClosePriceIndicator(series)

// record trades on this object
record := talib4g.NewTradingRecord()

entryConstant := talib4g.NewConstantIndicator(30)
exitConstant := talib4g.NewConstantIndicator(10)

// Is satisfied when the price ema moves above 30 and the current position is new
entryRule := talib4g.And(
	talib4g.NewCrossUpIndicatorRule(entryConstant, indicator),
	talib4g.NewPositionNewRule())
	
// Is satisfied when the price ema moves below 10 and the current position is open
exitRule := talib4g.And(
	talib4g.NewCrossDownIndicatorRule(indicator, exitConstant),
	talib4g.NewPositionOpenRule()) 

strategy := talib4g.RuleStrategy{
	UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
	EntryRule:      entryRule,
	ExitRule:       exitRule,
}

strategy.ShouldEnter(0, record) // returns false
```

### Credits
Talib4g is heavily influenced by the great [ta4j](https://github.com/ta4j/ta4j). Many of the ideas and frameworks in this library owe their genesis to the great work done over there.

### License

Talib4g is released under the MIT license. See [LICENSE](./LICENSE) for details.
