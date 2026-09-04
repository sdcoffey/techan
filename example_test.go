package techan_test

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func ExampleNewMACDHistogramIndicator() {
	prices := techan.NewFixedIndicator(100, 100, 100, 100, 100, 100)
	macd := techan.NewMACDIndicator(prices, 2, 4)
	histogram := techan.NewMACDHistogramIndicator(macd, 2)
	first := techan.FirstValidIndex(histogram)
	fmt.Println(first, histogram.Calculate(first))
	// Output: 4 0
}

func ExampleResetCacheFrom() {
	series := techan.NewTimeSeries()
	for i := 0; i < 4; i++ {
		candle := techan.NewCandle(techan.NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.ClosePrice = big.NewFromInt(10)
		series.AddCandle(candle)
	}
	macd := techan.NewMACDIndicator(techan.NewClosePriceIndicator(series), 2, 3)
	fmt.Println(macd.Calculate(3).FormattedString(4))
	series.Candles[3].ClosePrice = big.NewFromInt(20)
	techan.ResetCacheFrom(macd, 3)
	fmt.Println(macd.Calculate(3).FormattedString(4))
	// Output:
	// 0.0000
	// 1.6667
}

func ExampleRuleStrategy() {
	prices := techan.NewFixedIndicator(99, 101, 102, 98)
	threshold := techan.NewConstantIndicator(100)
	strategy := techan.RuleStrategy{
		EntryRule: techan.NewCrossUpIndicatorRule(threshold, prices),
		ExitRule:  techan.NewCrossDownIndicatorRule(prices, threshold),
	}
	record := techan.NewTradingRecord()
	for i := 0; i < 4; i++ {
		if strategy.ShouldEnter(i, record) {
			record.Operate(techan.Order{Side: techan.BUY, Price: prices.Calculate(i), Amount: big.ONE, ExecutionTime: time.Unix(int64(i), 0)})
			fmt.Println("enter", i)
		} else if strategy.ShouldExit(i, record) {
			record.Operate(techan.Order{Side: techan.SELL, Price: prices.Calculate(i), Amount: big.ONE, ExecutionTime: time.Unix(int64(i), 0)})
			fmt.Println("exit", i)
		}
	}
	fmt.Println("profit", techan.TotalProfitAnalysis{}.Analyze(record))
	// Output:
	// enter 1
	// exit 3
	// profit -3
}
