package talib4g

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

type Analysis interface {
	Analyze(*TradingRecord) decimal.Decimal
}

type TotalProfitAnalysis string

func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) decimal.Decimal {
	profit := decimal.NewFromFloat(0)
	for _, trade := range record.Trades {
		costBasis := trade.EntranceOrder().Amount.Mul(trade.EntranceOrder().Price)
		sellPrice := trade.ExitOrder().Amount.Mul(trade.ExitOrder().Price)

		profit = profit.Add(sellPrice.Sub(costBasis))
	}

	return profit
}

type NumTradesAnalysis string

func (nta NumTradesAnalysis) Analyze(record *TradingRecord) decimal.Decimal {
	return decimal.New(int64(len(record.Trades)), 0)
}

type LogTradesAnalysis string

func (lta LogTradesAnalysis) Analyze(record *TradingRecord) decimal.Decimal {
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

		log.Println(fmt.Sprintf("%s - %s with %s (%s @ $%s)", order.ExecutionTime.Format(time.RFC822), action, oType, order.Amount, order.Price))
	}

	for _, trade := range record.Trades {
		logOrder(trade.EntranceOrder())
		logOrder(trade.ExitOrder())
	}
	return decimal.Zero
}
