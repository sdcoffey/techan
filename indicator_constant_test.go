package techan

import (
	"math"
	"testing"
)

func TestConstantIndicator_Calculate(t *testing.T) {
	ci := NewConstantIndicator(4.56)

	decimalEquals(t, 4.56, ci.Calculate(0))
	decimalEquals(t, 4.56, ci.Calculate(-math.MaxInt64))
	decimalEquals(t, 4.56, ci.Calculate(math.MaxInt64))
}
