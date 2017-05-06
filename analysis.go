package talib4g

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Analysis interface {
	Analyze(*TradingRecord) float64
}

type TotalProfitAnalysis float64

func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	profit := 0.0
	for _, trade := range record.Trades {
		costBasis := trade.EntranceOrder().Amount * trade.EntranceOrder().Price
		costBasis *= float64(1 + tps)
		sellPrice := trade.ExitOrder().Amount * trade.ExitOrder().Price

		profit += sellPrice - costBasis
	}

	return profit
}

type NumTradesAnalysis string

func (nta NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	return float64(len(record.Trades))
}

type LogTradesAnalysis struct {
	io.Writer
}

func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	logOrder := func(order *Order) {
		var oType string
		var action string
		if order.Type == BUY {
			oType = "buy"
			action = "Entered"
		} else {
			oType = "sell"
			action = "Exited"
		}

		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - %s with %s (%f @ $%f)", order.ExecutionTime.Format(time.RFC822), action, oType, order.Amount, order.Price))
	}

	for _, trade := range record.Trades {
		logOrder(trade.EntranceOrder())
		logOrder(trade.ExitOrder())
	}
	return 0.0
}

type PeriodProfitAnalysis time.Duration

func (ppa PeriodProfitAnalysis) Analyze(record *TradingRecord) float64 {
	var tp TotalProfitAnalysis
	totalProfit := tp.Analyze(record)

	periods := record.Trades[len(record.Trades)-1].ExitOrder().ExecutionTime.Sub(record.Trades[0].EntranceOrder().ExecutionTime) / time.Duration(ppa)
	return totalProfit / float64(periods)
}

type ProfitableTradesAnalysis string

func (pta ProfitableTradesAnalysis) Analyze(record *TradingRecord) float64 {
	var profitableTrades int
	for _, trade := range record.Trades {
		costBasis := trade.EntranceOrder().Amount * trade.EntranceOrder().Price
		sellPrice := trade.ExitOrder().Amount * trade.ExitOrder().Price

		if sellPrice > costBasis {
			profitableTrades++
		}
	}

	return float64(profitableTrades)
}

type AverageProfitAnalysis string

func (apa AverageProfitAnalysis) Analyze(record *TradingRecord) float64 {
	var tp TotalProfitAnalysis
	totalProft := tp.Analyze(record)

	return totalProft / float64(len(record.Trades))
}

type graphAnalysis struct {
	output *os.File
	record TradingRecord
	series TimeSeries
}
