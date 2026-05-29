# Techan Release notes

## 0.13.0
* **BREAKING**: Raise the minimum supported Go version from 1.12 to 1.21.
* Replace Travis CI with GitHub Actions and test Go 1.21 through 1.26.
* Fix stop-loss percentage calculations for multi-unit orders and short positions (fixes #8).
* Fix trendline divide-by-zero behavior for sparse windows.
* Add cache invalidation support for built-in cached indicators.
* Document SMA and EMA warm-up behavior.
* Harden analysis helpers for empty records, invalid periods, zero cost basis, and short trades.
* Update Go module dependencies, including `github.com/sdcoffey/big` to v0.8.0 and `github.com/stretchr/testify` to v1.11.1.

## 0.12.1
* Fixes EMA window calculation (thanks @danhenke and @joelnordell!)

## 0.12.0
* Add MaximumValue and MinimumValue Indicators
* Add [MaximumDrawdownIndicator](https://www.investopedia.com/terms/m/maximum-drawdown-mdd.asp).

## 0.11.0
* Add BollingerUpperBandIndicator and BollingerLowerBandIndicator (thanks @shkim!)

## 0.10.0
* Add TimePeriod#In to modify timezone information
* Add TimePeriod#UTC to set location to UTC

## 0.9.0
* Add [AroonIndicator](https://www.investopedia.com/terms/a/aroon.asp)
* Deprecate [Parse](https://godoc.org/github.com/sdcoffey/techan#Parse) for time periods, introduce `ParseTimePeriod`

## 0.8.0
* Add [MMAIndicator](https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average)
* Add GainIndicator
* Add LossIndicator
* Fix bug in RSI calculation (fixes #13)

## 0.7.1
* Fix bug in trendline indicator to prevent low-end index-OOB errors

## 0.7.0
* Add Trendline indicator
* Update big to v0.4.1

## 0.6.1
* Merge #10, which fixes a bug in TotalProfitAnalysis not taking short positions into effect

## 0.6.0
* **BREAKING**: Standard Deviation Indicator and Variance indicator now use the NewXIndicator pattern used throughout the library. Any usages creating the struct directly will need to be udpated.
* Migrate to go module

## 0.5.0
* Add StandardDeviationIndicator
* Add VarianceIndicator

## 0.4.0
* Add DerivativeIndicator

## 0.3.0
* Rename talib4g to techan

## 0.2.0
* Remove NewOrder methods and prefer struct initializer
* Add missing test coverage
* Add godoc

## 0.1.1
* Update documentation

## 0.1.0
* Initial release of talib4g
* Support for basic indicators
* Support for timeseries
* Support for basic strategies
* Support for entry and exit rules
