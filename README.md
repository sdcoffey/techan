## Techan
[![Tests](https://github.com/sdcoffey/techan/actions/workflows/test.yml/badge.svg)](https://github.com/sdcoffey/techan/actions/workflows/test.yml)

[![codecov](https://codecov.io/gh/sdcoffey/techan/branch/master/graph/badge.svg)](https://codecov.io/gh/sdcoffey/techan)

TechAn is a **tech**nical **an**alysis library for Go! It provides a suite of tools and frameworks to analyze financial data and make trading decisions.

## Features 
* Basic and advanced technical analysis indicators
* Profit and trade analysis
* Strategy building

Requires **Go 1.21 or newer**.

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

if series.LastIndex() >= techan.FirstValidIndex(movingAverage) {
	fmt.Println(movingAverage.Calculate(series.LastIndex()).FormattedString(2))
}
```

Windowed indicators such as SMA and EMA return `0` until enough data exists to fill their first window. For example, a 10-period EMA starts producing calculated values at index `9`.

Use `techan.FirstValidIndex(indicator)` when composing indicators. For example,
MACD(12,26) first becomes usable at index 25, and its 9-period signal at index 33.
Custom indicators can implement `WarmupIndicator` to declare their first usable
index; otherwise they are assumed ready at zero. RSI uses Wilder's seed of N
actual price changes, so RSI(14) starts at index 14; a flat window returns zero.
Bollinger bands and windowed standard deviation return zero until a complete
window is available.

EMA and MMA fill caches iteratively and round stored values to 256 binary bits
(about 77 significant decimal digits). This keeps memory proportional to history
length without narrowing prices to `float64`. Caches still retain historical
results. After editing candles, call `techan.ResetCacheFrom(rootIndicator, index)`
on each indicator graph you use; it reaches built-in nested caches. For a
reindexed or truncated series, reset from zero or rebuild the indicators. A full
reset recomputes from the retained history; it does not preserve EMA history that
has been removed. Mutation, calculation, and resetting must be serialized by the
caller when indicators are shared between goroutines.

### Creating trading strategies
```go
indicator := techan.NewClosePriceIndicator(series)

// record trades on this object
record := techan.NewTradingRecord()

entryConstant := techan.NewConstantIndicator(30)
exitConstant := techan.NewConstantIndicator(10)

// Is satisfied when the close price moves above 30 and the current position is new
entryRule := techan.And(
	techan.NewCrossUpIndicatorRule(entryConstant, indicator),
	techan.PositionNewRule{})
	
// Is satisfied when the close price moves below 10 and the current position is open
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

Crossover rules fire only on a crossing bar. Equality alone does not trigger a
crossing; an equality plateau is resolved using the preceding unequal values.
`UnstablePeriod` is the first index where the strategy may enter or exit.

### Analysis conventions

Profit analyses account for long and short order direction. `PercentGainAnalysis`
returns total realized profit divided by the first closed position's entry cost,
which is treated as initial capital. It does not model deposits, withdrawals, or
fees. `PeriodProfitAnalysis` includes fractional periods. The buy-and-hold
comparison uses only completed candles through the record's last closed trade.
Maximum drawdown measures declines after positive peaks within the requested
window; a window without a positive peak returns zero.

### Contributing

- `make test` runs all tests without rewriting source files.
- `make format` applies goimports; `make lint` runs vet and Staticcheck.
- `make bench` reports timing and allocations, including EMA/MMA history scaling.
- `go test -run '^$' -fuzz FuzzParseTimePeriod -fuzztime 10s` exercises parsing.

Formatting and lint tools are version-pinned in the Makefile and run separately
from the library module. They may automatically select a newer Go toolchain;
the library's minimum remains Go 1.21. See the executable examples in the test
files for indicator composition, cache resets, and a complete trading record.

### Enjoying this project?
Are you using techan in production? You can sponsor its development by buying me a coffee! ☕

**ETH:** `0x2D9d3A1c16F118A3a59d0e446d574e1F01F62949`

### Credits
Techan is heavily influenced by the great [ta4j](https://github.com/ta4j/ta4j). Many of the ideas and frameworks in this library owe their genesis to the great work done over there.

### License

Techan is released under the MIT license. See [LICENSE](./LICENSE) for details.
