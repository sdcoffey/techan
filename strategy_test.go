package techan

import (
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

type alwaysSatisfiedRule struct{}

func (a alwaysSatisfiedRule) IsSatisfied(index int, record *TradingRecord) bool {
	return true
}

func TestRuleStrategy_ShouldEnter(t *testing.T) {
	t.Run("Returns false if index < unstable period", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.False(t, s.ShouldEnter(0, record))
	})

	t.Run("Returns false if a position is open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.ONE,
		})

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.False(t, s.ShouldEnter(6, record))
	})

	t.Run("Returns true when position is closed", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.True(t, s.ShouldEnter(6, record))
	})

	t.Run("panics when entry rule is nil", func(t *testing.T) {
		s := RuleStrategy{
			ExitRule:       alwaysSatisfiedRule{},
			UnstablePeriod: 10,
		}

		assert.PanicsWithValue(t, "entry rule cannot be nil", func() {
			s.ShouldEnter(0, nil)
		})
	})
}

func TestRuleStrategy_ShouldExit(t *testing.T) {
	t.Run("Returns false if index < unstablePeriod", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.ONE,
		})

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.False(t, s.ShouldExit(0, record))
	})

	t.Run("Returns false when position is closed", func(t *testing.T) {
		record := NewTradingRecord()

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.False(t, s.ShouldExit(6, record))
	})

	t.Run("Returns true when position is open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: big.ONE,
			Price:  big.ONE,
		})

		s := RuleStrategy{
			alwaysSatisfiedRule{},
			alwaysSatisfiedRule{},
			5,
		}

		assert.True(t, s.ShouldExit(6, record))
	})

	t.Run("panics when exit rule is nil", func(t *testing.T) {
		s := RuleStrategy{
			EntryRule:      alwaysSatisfiedRule{},
			UnstablePeriod: 10,
		}

		assert.PanicsWithValue(t, "exit rule cannot be nil", func() {
			s.ShouldExit(0, nil)
		})
	})
}
