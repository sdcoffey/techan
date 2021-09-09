## Techan
![](https://travis-ci.org/sdcoffey/techan.svg?branch=master)

[![codecov](https://codecov.io/gh/sdcoffey/techan/branch/master/graph/badge.svg)](https://codecov.io/gh/sdcoffey/techan)

TechAn is a **tech**nical **an**alysis library for Go! It provides a suite of tools and frameworks to analyze financial data and make trading decisions.

## Features 
* Basic and advanced technical analysis indicators
* Profit and trade analysis
* Strategy building

### Installation
```sh
$ go get github.com/sdcoffey/techan
```

### Quickstart
```go
series := techan.NewTimeSeries()

// fetch this from your preferred exchange
dataset := [][]string{
	// Timestamp, Open, Close, High, Low, volume
	{"1234567", "1", "2", "3", "5", "6"},
}

for _, datum := range dataset {
	start, _ := strconv.ParseInt(datum[0], 10, 64)
	period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

	candle := techan.NewCandle(period)
	candle.OpenPrice = big.NewFromString(datum[1])
	candle.ClosePrice = big.NewFromString(datum[2])
	candle.MaxPrice = big.NewFromString(datum[3])
	candle.MinPrice = big.NewFromString(datum[4])

	series.AddCandle(candle)
}

closePrices := techan.NewClosePriceIndicator(series)
movingAverage := techan.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

fmt.Println(movingAverage.Calculate(0).FormattedString(2))
```

### Creating trading strategies
```go
indicator := techan.NewClosePriceIndicator(series)

// record trades on this object
record := techan.NewTradingRecord()

entryConstant := techan.NewConstantIndicator(30)
exitConstant := techan.NewConstantIndicator(10)

// Is satisfied when the price ema moves above 30 and the current position is new
entryRule := techan.And(
	techan.NewCrossUpIndicatorRule(entryConstant, indicator),
	techan.PositionNewRule{})
	
// Is satisfied when the price ema moves below 10 and the current position is open
exitRule := techan.And(
	techan.NewCrossDownIndicatorRule(indicator, exitConstant),
	techan.PositionOpenRule{})

strategy := techan.RuleStrategy{
	UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
	EntryRule:      entryRule,
	ExitRule:       exitRule,
}

strategy.ShouldEnter(0, record) // returns false
```

### Enjoying this project?
Are you using techan in production? You can sponsor its development by buying me a coffee! ☕

**BTC:** `1DipyghRPG8iFQzLPKMNEVA4gnJ8q8qQfH`

**ETH:** `0x4b86389f57da7ece64882b1c1f2e4ba698f5659a`

### Credits
Techan is heavily influenced by the great [ta4j](https://github.com/ta4j/ta4j). Many of the ideas and frameworks in this library owe their genesis to the great work done over there.

### License

Techan is released under the MIT license. See [LICENSE](./LICENSE) for details.
