package talib4g

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestConstantIndicator_Calculate(t *testing.T) {
	ci := NewConstantIndicator(4.56)

	assert.EqualValues(t, 4.56, ci.Calculate(0))
	assert.EqualValues(t, 4.56, ci.Calculate(-math.MaxInt64))
	assert.EqualValues(t, 4.56, ci.Calculate(math.MaxInt64))
}
