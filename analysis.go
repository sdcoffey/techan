package talib4g

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/sdcoffey/big"
)

type Analysis interface {
	Analyze(*TradingRecord) float64
}

type TotalProfitAnalysis float64

func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	totalProfit := big.NewDecimal(0)
	for _, trade := range record.Trades {
		if trade.IsClosed() {
			costBasis := trade.CostBasis().Frac(1 + float64(tps))
			exitValue := trade.ExitValue().Frac(math.Abs(float64(tps) - 1))
			totalProfit = totalProfit.Add(exitValue.Sub(costBasis))
		}
	}

	return totalProfit.Float()
}

type PercentGainAnalysis struct{}

func (pga PercentGainAnalysis) Analyze(record *TradingRecord) float64 {
	if len(record.Trades) > 0 && record.Trades[0].IsClosed() {
		return (record.Trades[len(record.Trades)-1].ExitValue().Div(record.Trades[0].CostBasis())).Sub(big.NewDecimal(1)).Float()
	} else {
		return 0
	}
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
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - enter with buy %s (%s @ $%s)", trade.EntranceOrder().ExecutionTime.UTC().Format(time.RFC822), trade.EntranceOrder().Security, trade.EntranceOrder().Amount, trade.EntranceOrder().Price))
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - exit with sell %s (%s @ $%s)", trade.ExitOrder().ExecutionTime.UTC().Format(time.RFC822), trade.ExitOrder().Security, trade.ExitOrder().Amount, trade.ExitOrder().Price))

		profit := trade.ExitValue().Sub(trade.CostBasis())
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
		costBasis := trade.EntranceOrder().Amount.Mul(trade.EntranceOrder().Price)
		sellPrice := trade.ExitOrder().Amount.Mul(trade.ExitOrder().Price)

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

type BuyAndHoldAnalysis struct {
	*TimeSeries
	StartingMoney float64
}

func (baha BuyAndHoldAnalysis) Analyze(record *TradingRecord) float64 {
	if len(record.Trades) == 0 {
		return 0
	}

	openOrder := NewOrder(BUY)
	openOrder.Amount = big.NewDecimal(baha.StartingMoney).Div(baha.Candles[0].ClosePrice)
	openOrder.Price = baha.Candles[0].ClosePrice

	closeOrder := NewOrder(SELL)
	closeOrder.Amount = openOrder.Amount
	closeOrder.Price = baha.Candles[len(baha.Candles)-1].ClosePrice

	pos := NewPosition(openOrder)
	pos.Exit(closeOrder)

	return pos.ExitValue().Sub(pos.CostBasis()).Float()
}
