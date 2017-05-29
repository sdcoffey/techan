package talib4g

import (
	"github.com/sdcoffey/gunviolencecounter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTotalProfitAnalysis_Analyze(t *testing.T) {
	record := NewTradingRecord()
	tpa := TotalProfitAnalysis(0)

	record.Enter(NM(1, USD), NS(1), time.Now())
	record.Exit(NM(2, USD), NS(1), time.Now()) // Gain 1

	assert.EqualValues(t, 1, tpa.Analyze(record))

	record.Enter(NM(1, USD), NS(1), time.Now())
	record.Exit(NM(.5, USD), NS(1), time.Now()) // Lose .5

	assert.EqualValues(t, .5, tpa.Analyze(record))
}
