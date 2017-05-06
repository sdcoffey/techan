package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type alwaysSatisfiedRule int

func (a alwaysSatisfiedRule) IsSatisfied(index int, record *TradingRecord) bool {
	return true
}

func TestRuleStrategy_ShouldEnter(t *testing.T) {
	t.Run("Returns false if index < unstable period", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.False(t, s.ShouldEnter(0, record))
	})

	t.Run("Returns false if a position is open", func(t *testing.T) {
		record := NewTradingRecord()
		record.Enter(0, 0, time.Now())

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.False(t, s.ShouldEnter(6, record))
	})

	t.Run("Returns true when position is closed", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.True(t, s.ShouldEnter(6, record))
	})
}

func TestRuleStrategy_ShouldExit(t *testing.T) {
	t.Run("Returns false if index < unstablePeriod", func(t *testing.T) {
		record := NewTradingRecord()
		record.Enter(0, 0, time.Now())

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.False(t, s.ShouldExit(0, record))
	})

	t.Run("Returns false when position is closed", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.False(t, s.ShouldExit(6, record))
	})

	t.Run("Returns true when position is open", func(t *testing.T) {
		record := NewTradingRecord()
		record.Enter(0, 0, time.Now())

		s := RuleStrategy{
			alwaysSatisfiedRule(0),
			alwaysSatisfiedRule(0),
			5,
		}

		assert.True(t, s.ShouldExit(6, record))
	})
}
