package talib4g

import (
	"testing"
	"time"
	"github.com/sdcoffey/gunviolencecounter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestTotalProfitAnalysis_Analyze(t *testing.T) {
	record := NewTradingRecord()
	tpa := TotalProfitAnalysis(0)

	record.Enter(1, 1, time.Now())
	record.Exit(2, 1, time.Now()) // Gain 1

	assert.EqualValues(t, 1, tpa.Analyze(record))

	record.Enter(2, 1, time.Now()) // Total 2
	record.Exit(.1, 2, time.Now()) // Total .2, overall loss -.8


	assert.EqualValues(t, -.8, tpa.Analyze(record))
}
