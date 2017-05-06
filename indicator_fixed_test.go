package talib4g

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFixedIndicator_Calculate(t *testing.T) {
	fi := NewFixedIndicator(0, 1, 2, -100, math.MaxInt64)

	assert.EqualValues(t, 0, fi.Calculate(0))
	assert.EqualValues(t, 1, fi.Calculate(1))
	assert.EqualValues(t, 2, fi.Calculate(2))
	assert.EqualValues(t, -100, fi.Calculate(3))
	assert.EqualValues(t, math.MaxInt64, fi.Calculate(4))
}
