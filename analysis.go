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
	profit := NM(0, USD)
	for _, trade := range record.Trades {
		if trade.IsClosed() {
			profit = profit.A(trade.SellValue().S(trade.CostBasis()))
		}
	}

	return profit.Float()
}

type NumTradesAnalysis string

func (nta NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	return float64(len(record.Trades))
}

type LogTradesAnalysis struct {
	io.Writer
}

func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	logOrder := func(trade *Position) {
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - enter with buy (%f @ $%f)", trade.EntranceOrder().ExecutionTime.UTC().Format(time.RFC822), trade.EntranceOrder().Amount, trade.EntranceOrder().Price))
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - exit with sell (%f @ $%f)", trade.ExitOrder().ExecutionTime.UTC().Format(time.RFC822), trade.ExitOrder().Amount, trade.ExitOrder().Price))

		profit := trade.ExitOrder().Amount.Convert(trade.ExitOrder().Price).S(trade.EntranceOrder().Amount.Convert(trade.EntranceOrder().Price))
		fmt.Fprintln(lta.Writer, fmt.Sprintf("Profit: $%.2f", profit))
	}

	for _, trade := range record.Trades {
		if trade.IsClosed() {
			logOrder(trade)
		}
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
		costBasis := trade.EntranceOrder().Amount.Convert(trade.EntranceOrder().Price)
		sellPrice := trade.ExitOrder().Amount.Convert(trade.ExitOrder().Price)

		if sellPrice.GT(costBasis) {
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
