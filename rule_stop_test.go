package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopLossRule(t *testing.T) {
	t.Run("Returns false when position is new or closed", func(t *testing.T) {
		record := NewTradingRecord()

		series := MockTimeSeries(1, 2, 3, 4)

		slr := NewStopLossRule(series, -0.1)

		assert.False(t, slr.IsSatisfied(3, record))
	})

	t.Run("Returns true when losses exceed tolerance", func(t *testing.T) {
		record := NewTradingRecord()
		record.Enter(NM(10, USD), NM(1, BTC), time.Now())

		series := MockTimeSeries(10, 9) // Lose 10%

		slr := NewStopLossRule(series, -0.05)

		assert.True(t, slr.IsSatisfied(1, record))
	})

	t.Run("Returns false when losses do not exceed tolerance", func(t *testing.T) {
		record := NewTradingRecord()
		record.Enter(NM(10, USD), NM(1, BTC), time.Now())

		series := MockTimeSeries(10, 10.1) // Lose 10%

		slr := NewStopLossRule(series, -0.05)

		assert.False(t, slr.IsSatisfied(1, record))
	})
}
