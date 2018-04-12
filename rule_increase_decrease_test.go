package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreaseRule(t *testing.T) {
	t.Run("returns false when index == 0", func(t *testing.T) {
		rule := IncreaseRule{}

		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("returns true when increase", func(t *testing.T) {
		series := mockTimeSeries("1", "2")
		rule := IncreaseRule{NewClosePriceIndicator(series)}

		assert.True(t, rule.IsSatisfied(1, nil))
	})

	t.Run("returns false when same", func(t *testing.T) {
		series := mockTimeSeries("1", "1")
		rule := IncreaseRule{NewClosePriceIndicator(series)}

		assert.False(t, rule.IsSatisfied(1, nil))
	})

	t.Run("returns false when decrease", func(t *testing.T) {
		series := mockTimeSeries("1", "0")
		rule := IncreaseRule{NewClosePriceIndicator(series)}

		assert.False(t, rule.IsSatisfied(1, nil))
	})
}

func TestDecreaseRule(t *testing.T) {
	t.Run("returns false when index == 0", func(t *testing.T) {
		rule := DecreaseRule{}

		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("returns true when decrease", func(t *testing.T) {
		series := mockTimeSeries("1", "0")
		rule := DecreaseRule{NewClosePriceIndicator(series)}

		assert.True(t, rule.IsSatisfied(1, nil))
	})

	t.Run("returns false when  decrease", func(t *testing.T) {
		series := mockTimeSeries("1", "2")
		rule := DecreaseRule{NewClosePriceIndicator(series)}

		assert.False(t, rule.IsSatisfied(1, nil))
	})

	t.Run("returns false when same", func(t *testing.T) {
		series := mockTimeSeries("1", "1")
		rule := IncreaseRule{NewClosePriceIndicator(series)}

		assert.False(t, rule.IsSatisfied(1, nil))
	})
}
