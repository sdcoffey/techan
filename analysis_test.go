package talib4g

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTotalProfitAnalysis_Analyze(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		record := NewTradingRecord()
		tpa := TotalProfitAnalysis(0)

		record.Enter(NewDecimal(1), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(2), NewDecimal(1), time.Now()) // Gain 1

		assert.EqualValues(t, 1, tpa.Analyze(record))

		record.Enter(NewDecimal(1), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(.5), NewDecimal(1), time.Now()) // Lose .5

		assert.EqualValues(t, .5, tpa.Analyze(record))
	})
}

func TestPercentGainAnalysis(t *testing.T) {
	t.Run("Simple gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NewDecimal(1), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(2), NewDecimal(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, 1, gain)
	})

	t.Run("Simple loss", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NewDecimal(2), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(1), NewDecimal(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.5, gain)
	})

	t.Run("Small loss and gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NewDecimal(2), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(1), NewDecimal(1), time.Now())

		record.Enter(NewDecimal(1), NewDecimal(1), time.Now())
		record.Exit(NewDecimal(1.25), NewDecimal(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.375, gain)
	})
}
