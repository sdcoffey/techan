package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceIndicator_Calculate(t *testing.T) {
	di := NewDifferenceIndicator(NewFixedIndicator(10, 9, 8), NewFixedIndicator(8, 9, 10))

	assert.EqualValues(t, 2, di.Calculate(0))
	assert.EqualValues(t, 0, di.Calculate(1))
	assert.EqualValues(t, -2, di.Calculate(2))
}
