package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTotalProfitAnalysis_Analyze(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		record := NewTradingRecord()
		tpa := TotalProfitAnalysis(0)

		record.Enter(NM(1, USD), NS(1), time.Now())
		record.Exit(NM(2, USD), NS(1), time.Now()) // Gain 1

		assert.EqualValues(t, 1, tpa.Analyze(record))

		record.Enter(NM(1, USD), NS(1), time.Now())
		record.Exit(NM(.5, USD), NS(1), time.Now()) // Lose .5

		assert.EqualValues(t, .5, tpa.Analyze(record))
	})
}

func TestPercentGainAnalysis(t *testing.T) {
	t.Run("Simple gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NM(1, USD), NS(1), time.Now())
		record.Exit(NM(2, USD), NS(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, 1, gain)
	})

	t.Run("Simple loss", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NM(2, USD), NS(1), time.Now())
		record.Exit(NM(1, USD), NS(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.5, gain)
	})

	t.Run("Small loss and gain", func(t *testing.T) {
		record := NewTradingRecord()

		pga := PercentGainAnalysis{}

		record.Enter(NM(2, USD), NS(1), time.Now())
		record.Exit(NM(1, USD), NS(1), time.Now())

		record.Enter(NM(1, USD), NS(1), time.Now())
		record.Exit(NM(1.25, USD), NS(1), time.Now())

		gain := pga.Analyze(record)
		assert.EqualValues(t, -.375, gain)
	})
}
