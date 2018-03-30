package talib4g

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

const example = "EXM"

func TestTotalProfitAnalysis_Analyze(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		record := NewTradingRecord()
		tpa := TotalProfitAnalysis(0)

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now()) // Gain 1

		assert.EqualValues(t, 1, tpa.Analyze(record))

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(.5), big.NewDecimal(1), big.ZERO, example, time.Now()) // Lose .5

		assert.EqualValues(t, .5, tpa.Analyze(record))
	})
}

func TestPercentGainAnalysis(t *testing.T) {
	t.Run("Simple gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, 1, gain)
	})

	t.Run("Simple loss", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.5, gain)
	})

	t.Run("Small loss and gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(big.NewDecimal(2), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())

		record.Enter(big.NewDecimal(1), big.NewDecimal(1), big.ZERO, example, time.Now())
		record.Exit(big.NewDecimal(1.25), big.NewDecimal(1), big.ZERO, example, time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.375, gain)
	})
}
