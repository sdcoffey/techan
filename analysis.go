package talib4g

import (
	"fmt"
	"io"
	"os"
	"time"
	"math"
)

type Analysis interface {
	Analyze(*TradingRecord) float64
}

type TotalProfitAnalysis float64

func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	totalProfit := NM(0, USD)
	for _, trade := range record.Trades {
		if trade.IsClosed() {
			costBasis := trade.CostBasis().Frac(1 + float64(tps))
			exitValue := trade.ExitValue().Frac(math.Abs(float64(tps) -1))
			totalProfit = totalProfit.A(exitValue.S(costBasis))
		}
	}

	return totalProfit.Float()
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
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - enter with buy (%s @ $%s)", trade.EntranceOrder().ExecutionTime.UTC().Format(time.RFC822), trade.EntranceOrder().Amount, trade.EntranceOrder().Price))
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - exit with sell (%s @ $%s)", trade.ExitOrder().ExecutionTime.UTC().Format(time.RFC822), trade.ExitOrder().Amount, trade.ExitOrder().Price))

		profit := trade.ExitValue().S(trade.CostBasis())
		fmt.Fprintln(lta.Writer, fmt.Sprintf("Profit: $%s", profit))
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
