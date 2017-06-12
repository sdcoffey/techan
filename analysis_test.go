package talib4g

import (
	"github.com/sdcoffey/gunviolencecounter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
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

	t.Run("Less simple", func(t *testing.T) {
		record := NewTradingRecord()
		tpa := TotalProfitAnalysis(0.002)

		record.Enter(NM(2241.60,USD), NM(0.07070000, BTC), time.Now())
		record.Exit(NM(2420.19,USD), NM(0.07070000, BTC), time.Now())

		record.Enter(NM(2024.80,USD), NM(0.08359999, BTC), time.Now())
		record.Exit(NM(2157.30,USD), NM(0.08359999, BTC), time.Now())

		record.Enter(NM(2212.90,USD), NM(0.08069999, BTC), time.Now())
		record.Exit(NM(2273.80,USD), NM(0.08069999, BTC), time.Now())

		record.Enter(NM(2386.30,USD), NM(0.07640000, BTC), time.Now())
		record.Exit(NM(2431.19,USD), NM(0.07640000, BTC), time.Now())

		record.Enter(NM(2518.40,USD), NM(0.07330000, BTC), time.Now())
		record.Exit(NM(2496.00,USD), NM(0.07330000, BTC), time.Now())

		record.Enter(NM(2539.90,USD), NM(0.07190000, BTC), time.Now())
		record.Exit(NM(2767.10,USD), NM(0.07190000, BTC), time.Now())

		record.Enter(NM(2767.60,USD), NM(0.07099999, BTC), time.Now())
		record.Enter(NM(2819.5,USD), NM(0.07099999, BTC), time.Now())

		record.Enter(NM(2819.50,USD), NM(0.07000000, BTC), time.Now())
		record.Enter(NM(2820.9,USD), NM(0.07000000, BTC), time.Now())

		record.Enter(NM(2821.20,USD), NM(0.07000000, BTC), time.Now())
		record.Enter(NM(2810.10,USD), NM(0.07000000, BTC), time.Now())

		profit := tpa.Analyze(record)
		assert.True(t, profit < 0)
	})
}

