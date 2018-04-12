package example

import "github.com/sdcoffey/techan"

func StrategyExample() {
	indicator := BasicEma()

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
