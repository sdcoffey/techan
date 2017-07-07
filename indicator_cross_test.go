package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrossIndicator(t *testing.T) {
	fixed := NewFixedIndicator(12, 11, 10, 9, 11, 8, 7, 6)
	cross := NewCrossIndicator(fixed, NewConstantIndicator(10))

	assert.EqualValues(t, 0, cross.Calculate(0))
	assert.EqualValues(t, 0, cross.Calculate(1))
	assert.EqualValues(t, 1, cross.Calculate(2))
	assert.EqualValues(t, 0, cross.Calculate(3))
	assert.EqualValues(t, 0, cross.Calculate(4))
	assert.EqualValues(t, 1, cross.Calculate(5))
	assert.EqualValues(t, 0, cross.Calculate(6))
	assert.EqualValues(t, 0, cross.Calculate(7))
}
