package example

import "github.com/sdcoffey/techan"

// StrategyExample shows how to create a simple trading strategy. In this example, a position should
// be opened if the price moves above 70, and the position should be closed if a position moves below 30.
func StrategyExample() {
	indicator := BasicEma() // from basic.go

	// record trades on this object
	record := techan.NewTradingRecord()

	entryConstant := techan.NewConstantIndicator(30)
	exitConstant := techan.NewConstantIndicator(10)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(entryConstant, indicator),
		techan.PositionNewRule{}) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(indicator, exitConstant),
		techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := techan.RuleStrategy{
		UnstablePeriod: 10,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	strategy.ShouldEnter(0, record) // returns false
}
