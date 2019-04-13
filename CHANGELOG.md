# Techan Release notes

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
