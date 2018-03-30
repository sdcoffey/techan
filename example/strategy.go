package example

import "github.com/sdcoffey/talib4g"

func StrategyExample() {
	indicator := BasicEma()

	// record trades on this object
	record := talib4g.NewTradingRecord()

	entryConstant := talib4g.NewConstantIndicator(30)
	exitConstant := talib4g.NewConstantIndicator(10)

	entryRule := talib4g.And(
		talib4g.NewCrossUpIndicatorRule(entryConstant, indicator),
		talib4g.NewPositionNewRule()) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := talib4g.And(
		talib4g.NewCrossDownIndicatorRule(indicator, exitConstant),
		talib4g.NewPositionOpenRule()) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := talib4g.RuleStrategy{
		UnstablePeriod: 10,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	strategy.ShouldEnter(0, record) // returns false
}
