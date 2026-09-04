package techan

import (
	"fmt"
	"io"
	"time"

	"github.com/sdcoffey/big"
)

// Analysis measures the performance of a trading record.
type Analysis interface {
	Analyze(*TradingRecord) float64
}

func positionProfit(trade *Position) big.Decimal {
	if trade == nil || !trade.IsClosed() {
		return big.ZERO
	}
	if trade.IsLong() {
		return trade.ExitValue().Sub(trade.CostBasis())
	}
	if trade.IsShort() {
		return trade.CostBasis().Sub(trade.ExitValue())
	}
	return big.ZERO
}

func totalProfit(record *TradingRecord) big.Decimal {
	total := big.ZERO
	if record != nil {
		for _, trade := range record.Trades {
			total = total.Add(positionProfit(trade))
		}
	}
	return total
}

func tradeBounds(record *TradingRecord) (first, last *Position) {
	if record == nil {
		return
	}
	for _, trade := range record.Trades {
		if trade == nil || !trade.IsClosed() {
			continue
		}
		if first == nil || trade.EntranceOrder().ExecutionTime.Before(first.EntranceOrder().ExecutionTime) {
			first = trade
		}
		if last == nil || trade.ExitOrder().ExecutionTime.After(last.ExitOrder().ExecutionTime) {
			last = trade
		}
	}
	return
}

// TotalProfitAnalysis measures realized profit, accounting for order direction.
type TotalProfitAnalysis struct{}

// Analyze returns total realized profit across closed positions.
func (TotalProfitAnalysis) Analyze(record *TradingRecord) float64 { return totalProfit(record).Float() }

// PercentGainAnalysis normalizes total realized profit by the first closed
// position's entry cost. Returns are fractions (0.1 means 10%). It assumes that
// this cost represents starting capital; deposits, withdrawals, and fees are not
// modeled. Both long and short positions contribute their directional profit.
type PercentGainAnalysis struct{}

// Analyze returns profit relative to starting capital, or zero for no closed
// positions or a zero initial cost basis.
func (PercentGainAnalysis) Analyze(record *TradingRecord) float64 {
	first, _ := tradeBounds(record)
	if first == nil || first.CostBasis().IsZero() {
		return 0
	}
	return totalProfit(record).Div(first.CostBasis()).Float()
}

// NumTradesAnalysis counts the trades in a record.
type NumTradesAnalysis string

// Analyze returns the number of recorded trades.
func (NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	if record == nil {
		return 0
	}
	return float64(len(record.Trades))
}

// LogTradesAnalysis writes closed positions and their realized profits.
type LogTradesAnalysis struct {
	io.Writer
}

// Analyze writes each closed position using its actual order sides.
func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	if record == nil {
		return 0
	}
	sideName := func(side OrderSide) string {
		switch side {
		case BUY:
			return "buy"
		case SELL:
			return "sell"
		default:
			return "unknown"
		}
	}
	for _, trade := range record.Trades {
		if trade == nil || !trade.IsClosed() {
			continue
		}
		entry, exit := trade.EntranceOrder(), trade.ExitOrder()
		fmt.Fprintf(lta.Writer, "%s - enter with %s %s (%s @ $%s)\n", entry.ExecutionTime.UTC().Format(time.RFC822), sideName(entry.Side), entry.Security, entry.Amount, entry.Price)
		fmt.Fprintf(lta.Writer, "%s - exit with %s %s (%s @ $%s)\n", exit.ExecutionTime.UTC().Format(time.RFC822), sideName(exit.Side), exit.Security, exit.Amount, exit.Price)
		fmt.Fprintf(lta.Writer, "Profit: $%s\n", positionProfit(trade))
	}
	return 0
}

// PeriodProfitAnalysis normalizes realized profit to a duration, including
// fractional periods. For example, profit of 15 over 90 minutes is 10 per hour.
type PeriodProfitAnalysis struct {
	Period time.Duration
}

// Analyze returns zero if there are no closed trades or the durations are invalid.
func (ppa PeriodProfitAnalysis) Analyze(record *TradingRecord) float64 {
	first, last := tradeBounds(record)
	if first == nil || ppa.Period <= 0 {
		return 0
	}
	elapsed := last.ExitOrder().ExecutionTime.Sub(first.EntranceOrder().ExecutionTime)
	if elapsed <= 0 {
		return 0
	}
	periods := float64(elapsed) / float64(ppa.Period)
	return totalProfit(record).Float() / periods
}

// ProfitableTradesAnalysis counts profitable closed positions.
type ProfitableTradesAnalysis struct{}

// Analyze counts profitable long and short positions.
func (ProfitableTradesAnalysis) Analyze(record *TradingRecord) float64 {
	if record == nil {
		return 0
	}
	count := 0
	for _, trade := range record.Trades {
		if positionProfit(trade).GT(big.ZERO) {
			count++
		}
	}
	return float64(count)
}

// AverageProfitAnalysis measures average realized profit per closed position.
type AverageProfitAnalysis struct{}

// Analyze returns average realized profit, or zero if no positions are closed.
func (AverageProfitAnalysis) Analyze(record *TradingRecord) float64 {
	if record == nil {
		return 0
	}
	count := 0
	for _, trade := range record.Trades {
		if trade != nil && trade.IsClosed() {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return totalProfit(record).Div(big.NewFromInt(count)).Float()
}

// BuyAndHoldAnalysis compares buying at the first available candle close with
// holding until the last closed trade's exit time. Only candles ending at or
// before that time are eligible, so future or unfinished bars cannot affect it.
type BuyAndHoldAnalysis struct {
	TimeSeries    *TimeSeries
	StartingMoney float64
}

// Analyze returns the hypothetical profit, or zero if no comparison is possible.
func (baha BuyAndHoldAnalysis) Analyze(record *TradingRecord) float64 {
	_, last := tradeBounds(record)
	if last == nil || baha.TimeSeries == nil || len(baha.TimeSeries.Candles) == 0 {
		return 0
	}
	firstClose := baha.TimeSeries.Candles[0].ClosePrice
	if firstClose.IsZero() {
		return 0
	}
	end := last.ExitOrder().ExecutionTime
	for i := len(baha.TimeSeries.Candles) - 1; i >= 0; i-- {
		candle := baha.TimeSeries.Candles[i]
		if !candle.Period.End.After(end) {
			amount := big.NewDecimal(baha.StartingMoney).Div(firstClose)
			return candle.ClosePrice.Sub(firstClose).Mul(amount).Float()
		}
	}
	return 0
}
